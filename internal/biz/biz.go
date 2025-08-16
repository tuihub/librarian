package biz

import (
	"github.com/tuihub/librarian/internal/biz/bizangela"
	"github.com/tuihub/librarian/internal/biz/bizbinah"
	"github.com/tuihub/librarian/internal/biz/bizchesed"
	"github.com/tuihub/librarian/internal/biz/bizgebura"
	"github.com/tuihub/librarian/internal/biz/bizkether"
	"github.com/tuihub/librarian/internal/biz/biznetzach"
	"github.com/tuihub/librarian/internal/biz/bizsupervisor"
	"github.com/tuihub/librarian/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/internal/biz/bizyesod"

	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	bizangela.NewAngela,
	bizkether.ProviderSet,
	biztiphereth.ProviderSet,
	bizgebura.NewGebura,
	bizbinah.ProviderSet,
	bizyesod.ProviderSet,
	biznetzach.ProviderSet,
	bizchesed.ProviderSet,
	bizsupervisor.ProviderSet,
)
