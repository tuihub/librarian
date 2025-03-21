package biz

import (
	"github.com/tuihub/librarian/internal/biz/bizbinah"
	"github.com/tuihub/librarian/internal/biz/bizchesed"
	"github.com/tuihub/librarian/internal/biz/bizgebura"
	"github.com/tuihub/librarian/internal/biz/bizkether"
	"github.com/tuihub/librarian/internal/biz/biznetzach"
	"github.com/tuihub/librarian/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/internal/biz/bizyesod"

	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	bizkether.ProviderSet,
	biztiphereth.ProviderSet,
	bizgebura.NewGebura,
	bizbinah.ProviderSet,
	bizyesod.ProviderSet,
	biznetzach.ProviderSet,
	bizchesed.ProviderSet,
)
