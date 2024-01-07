package biz

import (
	"github.com/google/wire"
	"github.com/tuihub/librarian/app/porter/internal/biz/bizs3"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(bizs3.NewS3)
