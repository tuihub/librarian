package bizchesed

import (
	"context"
	"strconv"
	"sync"

	"github.com/tuihub/librarian/internal/biz/bizutils"
	"github.com/tuihub/librarian/internal/data"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libidgenerator"
	"github.com/tuihub/librarian/internal/lib/libsearch"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelbinah"
	"github.com/tuihub/librarian/internal/model/modelchesed"
	miner "github.com/tuihub/protos/pkg/librarian/miner/v1"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewChesed,
	NewImageCache,
)

type Chesed struct {
	repo        *data.ChesedRepo
	b           *data.BinahRepo
	id          *libidgenerator.IDGenerator
	search      libsearch.Search
	porter      porter.LibrarianPorterServiceClient
	miner       miner.LibrarianMinerServiceClient
	upload      *modelbinah.UploadCallBack
	download    *modelbinah.DownloadCallBack
	imageCache  *libcache.Map[model.InternalID, modelchesed.Image]
	muScanImage sync.Mutex
}

func NewChesed(
	repo *data.ChesedRepo,
	b *data.BinahRepo,
	id *libidgenerator.IDGenerator,
	search libsearch.Search,
	cron *libcron.Cron,
	pClient porter.LibrarianPorterServiceClient,
	// miClient miner.LibrarianMinerServiceClient,
	block *modelbinah.ControlBlock,
	imageCache *libcache.Map[model.InternalID, modelchesed.Image],
) (*Chesed, error) {
	c := &Chesed{
		repo:        repo,
		b:           b,
		id:          id,
		search:      search,
		porter:      pClient,
		miner:       nil,
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
	err := cron.BySeconds("ChesedScanImage", 60, c.ScanImage, context.Background()) //nolint:mnd //TODO
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
	id, err := c.id.New()
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

func (c *Chesed) ScanImage(ctx context.Context) error {
	if c.muScanImage.TryLock() {
		defer c.muScanImage.Unlock()
	} else {
		return nil
	}
	images, err0 := c.repo.ListImageNeedScan(ctx)
	if err0 != nil {
		return err0
	}
	if len(images) == 0 {
		return nil
	}
	for _, image := range images {
		data, err := c.b.PresignedGetObject(
			ctx,
			data.BucketDefault,
			strconv.FormatInt(int64(image.ID), 10),
			libtime.Day,
		)
		if err != nil {
			return err
		}
		results, err := c.miner.RecognizeImageURL(ctx, &miner.RecognizeImageURLRequest{Url: data})
		if err != nil {
			return err
		}
		var desReq string
		for _, r := range results.GetResults() {
			desReq += r.GetText() + " "
		}
		if err = c.search.DescribeID(ctx,
			image.ID,
			libsearch.SearchIndexChesedImage,
			true,
			desReq,
		); err != nil {
			return err
		}
		if err = c.repo.SetImageStatus(ctx, image.ID, modelchesed.ImageStatusScanned); err != nil {
			return err
		}
	}
	return nil
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

func (c *Chesed) SearchImages(ctx context.Context, paging model.Paging, query string) (
	[]model.InternalID, *errors.Error) {
	if libauth.FromContextAssertUserType(ctx) == nil {
		return nil, bizutils.NoPermissionError()
	}
	results, err := c.search.SearchID(ctx,
		libsearch.SearchIndexChesedImage,
		paging,
		query,
	)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	ids := make([]model.InternalID, 0, len(results))
	for _, r := range results {
		ids = append(ids, r.ID)
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
