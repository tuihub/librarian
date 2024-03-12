package biz

import (
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizangela"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizbinah"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizchesed"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biznetzach"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizyesod"

	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	bizangela.ProviderSet,
	biztiphereth.ProviderSet,
	bizgebura.NewGebura,
	bizbinah.NewBinah,
	bizbinah.NewControlBlock,
	bizyesod.NewYesod,
	biznetzach.NewNetzach,
	bizchesed.ProviderSet,
)
