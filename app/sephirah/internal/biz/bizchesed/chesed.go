package bizchesed

import (
	"context"
	"strconv"
	"sync"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizbinah"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizutils"
	"github.com/tuihub/librarian/app/sephirah/internal/client"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelbinah"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelchesed"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/model"
	miner "github.com/tuihub/protos/pkg/librarian/miner/v1"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	searcherpb "github.com/tuihub/protos/pkg/librarian/searcher/v1"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewChesed,
	NewImageCache,
)

type ChesedRepo interface {
	CreateImage(context.Context, model.InternalID, *modelchesed.Image) error
	ListImages(context.Context, model.InternalID, model.Paging) ([]*modelchesed.Image, int64, error)
	ListImageNeedScan(context.Context) ([]*modelchesed.Image, error)
	SetImageStatus(context.Context, model.InternalID, modelchesed.ImageStatus) error
	GetImage(context.Context, model.InternalID, model.InternalID) (*modelchesed.Image, error)
}

type Chesed struct {
	repo        ChesedRepo
	b           bizbinah.BinahRepo
	searcher    *client.Searcher
	porter      porter.LibrarianPorterServiceClient
	miner       miner.LibrarianMinerServiceClient
	upload      *modelbinah.UploadCallBack
	download    *modelbinah.DownloadCallBack
	imageCache  *libcache.Map[model.InternalID, modelchesed.Image]
	muScanImage sync.Mutex
}

func NewChesed(
	repo ChesedRepo,
	b bizbinah.BinahRepo,
	cron *libcron.Cron,
	pClient porter.LibrarianPorterServiceClient,
	sClient *client.Searcher,
	miClient miner.LibrarianMinerServiceClient,
	block *modelbinah.ControlBlock,
	imageCache *libcache.Map[model.InternalID, modelchesed.Image],
) (*Chesed, error) {
	c := &Chesed{
		repo:        repo,
		b:           b,
		porter:      pClient,
		searcher:    sClient,
		miner:       miClient,
		upload:      nil,
		download:    nil,
		imageCache:  imageCache,
		muScanImage: sync.Mutex{},
	}
	c.upload = block.RegisterUploadCallback(
		modelbinah.UploadChesedImage,
		c.UploadImageCallback,
	)
	c.download = block.RegisterDownloadCallback(
		modelbinah.DownloadEmpty,
		nil,
	)
	err := cron.BySeconds(60, c.ScanImage, context.Background()) //nolint:gomnd //TODO
	if err != nil {
		return nil, err
	}
	return c, nil
}

func NewImageCache(
	store libcache.Store,
) *libcache.Map[model.InternalID, modelchesed.Image] {
	return libcache.NewMap[model.InternalID, modelchesed.Image](
		store,
		"Images",
		func(k model.InternalID) string {
			return strconv.FormatInt(int64(k), 10)
		},
		nil,
		libcache.WithExpiration(libtime.Day),
	)
}

func (c *Chesed) UploadImage(ctx context.Context, image modelchesed.Image,
	metadata modelbinah.FileMetadata) (string, *errors.Error) {
	if libauth.FromContextAssertUserType(ctx) == nil {
		return "", bizutils.NoPermissionError()
	}
	if err := metadata.Check(); err != nil {
		return "", pb.ErrorErrorReasonBadRequest("invalid file metadata: %s", err.Error())
	}
	id, err := c.searcher.NewID(ctx)
	if err != nil {
		return "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	image.ID = id
	metadata.ID = id
	err = c.imageCache.Set(ctx, image.ID, &image)
	if err != nil {
		return "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	token, err := c.upload.GenerateUploadToken(ctx, metadata, libtime.HalfDay)
	if err != nil {
		return "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return token, nil
}

func (c *Chesed) UploadImageCallback(ctx context.Context, id model.InternalID) error {
	claims := libauth.FromContext(ctx)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	image, err := c.imageCache.Get(ctx, id)
	if err != nil {
		return err
	}
	err = c.repo.CreateImage(ctx, claims.UserID, image)
	if err != nil {
		return err
	}
	return nil
}

func (c *Chesed) ScanImage(ctx context.Context) {
	if c.muScanImage.TryLock() {
		defer c.muScanImage.Unlock()
	} else {
		return
	}
	images, err0 := c.repo.ListImageNeedScan(ctx)
	if err0 != nil {
		return
	}
	if len(images) == 0 {
		return
	}
	for _, image := range images {
		data, err := c.b.PresignedGetObject(ctx, bizbinah.BucketDefault, strconv.FormatInt(int64(image.ID), 10), libtime.Day)
		if err != nil {
			return
		}
		results, err := c.miner.RecognizeImageURL(ctx, &miner.RecognizeImageURLRequest{Url: data})
		if err != nil {
			return
		}
		var desReq string
		for _, r := range results.GetResults() {
			desReq += r.GetText() + " "
		}
		if err = c.searcher.DescribeID(ctx,
			image.ID,
			desReq,
			searcherpb.DescribeIDRequest_DESCRIBE_MODE_APPEND,
			searcherpb.Index_INDEX_CHESED_IMAGE,
		); err != nil {
			return
		}
		if err = c.repo.SetImageStatus(ctx, image.ID, modelchesed.ImageStatusScanned); err != nil {
			return
		}
	}
}

func (c *Chesed) ListImages(ctx context.Context, paging model.Paging) ([]model.InternalID, int64, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, 0, bizutils.NoPermissionError()
	}
	images, total, err := c.repo.ListImages(ctx, claims.UserID, paging)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	res := make([]model.InternalID, 0, len(images))
	for _, image := range images {
		res = append(res, image.ID)
	}
	return res, total, nil
}

func (c *Chesed) SearchImages(ctx context.Context, paging model.Paging, keywords string) (
	[]model.InternalID, *errors.Error) {
	if libauth.FromContextAssertUserType(ctx) == nil {
		return nil, bizutils.NoPermissionError()
	}
	ids, err := c.searcher.SearchID(ctx,
		paging,
		keywords,
		searcherpb.Index_INDEX_CHESED_IMAGE,
	)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return ids, nil
}

func (c *Chesed) DownloadImage(ctx context.Context, id model.InternalID) (string, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return "", pb.ErrorErrorReasonUnauthorized("no permission")
	}
	image, err := c.repo.GetImage(ctx, claims.UserID, id)
	if err != nil {
		return "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	token, err := c.download.GenerateDownloadToken(ctx, modelbinah.FileMetadata{
		ID:        id,
		Name:      image.Name,
		SizeBytes: 0,
		Type:      0,
		Sha256:    nil,
	}, libtime.HalfDay)
	if err != nil {
		return "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return token, nil
}
