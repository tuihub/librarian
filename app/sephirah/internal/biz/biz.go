package biz

import (
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizangela"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizbinah"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"

	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	bizangela.ProviderSet,
	biztiphereth.NewTiphereth,
	bizgebura.NewGebura,
	bizbinah.NewBinah,
	bizbinah.NewCallbackControl,
)
