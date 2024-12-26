package bizbinah

import (
	"context"
	"io"
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modelbinah"
	"github.com/tuihub/librarian/internal/lib/libauth"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewBinah,
	NewControlBlock,
)

type Binah struct {
	repo     BinahRepo
	callback *modelbinah.ControlBlock
	auth     *libauth.Auth
}

type BinahRepo interface {
	FeatureEnabled() bool
	PutObject(context.Context, io.Reader, Bucket, string) error
	PresignedGetObject(context.Context, Bucket, string, time.Duration) (string, error)
	PresignedPutObject(context.Context, Bucket, string, time.Duration) (string, error)
}

type Bucket int

const (
	BucketUnspecified Bucket = iota
	BucketDefault
)

func NewBinah(
	repo BinahRepo,
	callback *modelbinah.ControlBlock,
	auth *libauth.Auth,
) *Binah {
	return &Binah{
		repo:     repo,
		callback: callback,
		auth:     auth,
	}
}

func NewControlBlock(a *libauth.Auth) *modelbinah.ControlBlock {
	return modelbinah.NewControlBlock(a)
}
