package modelbinah

import (
	"context"
	"errors"
	"time"

	"github.com/tuihub/librarian/internal/lib/libauth"
)

type UploadCallBack struct {
	id           UploadCallbackID
	controlBlock *ControlBlock
}

type DownloadCallBack struct {
	id           DownloadCallbackID
	controlBlock *ControlBlock
}

func (u *UploadCallBack) GenerateUploadToken(ctx context.Context, meta FileMetadata,
	expire time.Duration) (string, error) {
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return "", errors.New("token required")
	}
	return u.controlBlock.a.GenerateToken(
		claims.InternalID,
		libauth.ClaimsTypeUploadToken,
		claims.UserType,
		&uploadTokenPayload{
			meta,
			u.id,
		},
		expire,
	)
}

func (u *DownloadCallBack) GenerateDownloadToken(ctx context.Context, meta FileMetadata,
	expire time.Duration) (string, error) {
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return "", errors.New("token required")
	}
	return u.controlBlock.a.GenerateToken(
		claims.InternalID,
		libauth.ClaimsTypeDownloadToken,
		claims.UserType,
		&downloadTokenPayload{
			meta,
			u.id,
		},
		expire,
	)
}
