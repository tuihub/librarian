package bizchesed

import (
	"context"
	"strconv"

	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelbinah"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelchesed"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/internal/model"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewChesed,
	NewImageCache,
)

type ChesedRepo interface {
}

type Chesed struct {
	// repo       ChesedRepo
	mapper     mapper.LibrarianMapperServiceClient
	searcher   searcher.LibrarianSearcherServiceClient
	porter     porter.LibrarianPorterServiceClient
	upload     *modelbinah.UploadCallBack
	imageCache *libcache.Map[model.InternalID, modelchesed.Image]
}

func NewChesed(
	// repo ChesedRepo,
	mClient mapper.LibrarianMapperServiceClient,
	pClient porter.LibrarianPorterServiceClient,
	sClient searcher.LibrarianSearcherServiceClient,
	block *modelbinah.ControlBlock,
	imageCache *libcache.Map[model.InternalID, modelchesed.Image],
) (*Chesed, error) {
	upload := block.RegisterUploadCallback(
		modelbinah.UploadArtifacts,
		UploadImageCallback,
	)
	y := &Chesed{
		//repo:       repo,
		mapper:     mClient,
		porter:     pClient,
		searcher:   sClient,
		upload:     upload,
		imageCache: imageCache,
	}
	return y, nil
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
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return "", pb.ErrorErrorReasonForbidden("no permission")
	}
	if err := metadata.Check(); err != nil {
		return "", pb.ErrorErrorReasonBadRequest("invalid file metadata: %s", err.Error())
	}
	resp, err := c.searcher.NewID(ctx, &searcher.NewIDRequest{})
	if err != nil {
		return "", pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	image.ID = converter.ToBizInternalID(resp.Id)
	metadata.ID = converter.ToBizInternalID(resp.Id)
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

func UploadImageCallback() error {
	// TODO
	return nil
}
