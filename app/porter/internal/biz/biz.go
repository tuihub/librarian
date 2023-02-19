package biz

import (
	"github.com/tuihub/librarian/app/porter/internal/biz/bizfeed"
	"github.com/tuihub/librarian/app/porter/internal/biz/bizs3"
	"github.com/tuihub/librarian/app/porter/internal/biz/bizsteam"

	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(bizsteam.NewSteamUseCase, bizs3.NewS3, bizfeed.NewFeed)
