package service

import (
	"github.com/tuihub/librarian/app/porter/internal/biz/bizsteam"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

func toPBAppType(t bizsteam.AppType) librarian.AppType {
	switch t {
	case bizsteam.AppTypeGame:
		return librarian.AppType_APP_TYPE_GAME
	default:
		return librarian.AppType_APP_TYPE_UNSPECIFIED
	}
}
