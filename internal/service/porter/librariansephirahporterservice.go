package porter

import (
	"github.com/tuihub/librarian/internal/biz/bizbinah"
	"github.com/tuihub/librarian/internal/biz/bizchesed"
	"github.com/tuihub/librarian/internal/biz/bizgebura"
	"github.com/tuihub/librarian/internal/biz/bizkether"
	"github.com/tuihub/librarian/internal/biz/biznetzach"
	"github.com/tuihub/librarian/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/internal/biz/bizyesod"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/service/supervisor"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1/porter"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewLibrarianSephirahPorterService)

type LibrarianSephirahPorterService struct {
	pb.UnimplementedLibrarianSephirahPorterServiceServer

	a    *bizkether.Kether
	t    *biztiphereth.Tiphereth
	g    *bizgebura.Gebura
	b    *bizbinah.Binah
	y    *bizyesod.Yesod
	n    *biznetzach.Netzach
	c    *bizchesed.Chesed
	s    *supervisor.Supervisor
	app  *libapp.Settings
	auth *libauth.Auth
}

func NewLibrarianSephirahPorterService(
	a *bizkether.Kether,
	t *biztiphereth.Tiphereth,
	g *bizgebura.Gebura,
	b *bizbinah.Binah,
	y *bizyesod.Yesod,
	n *biznetzach.Netzach,
	c *bizchesed.Chesed,
	s *supervisor.Supervisor,
	app *libapp.Settings,
	auth *libauth.Auth,
) pb.LibrarianSephirahPorterServiceServer {
	res := &LibrarianSephirahPorterService{
		UnimplementedLibrarianSephirahPorterServiceServer: pb.UnimplementedLibrarianSephirahPorterServiceServer{},
		a:    a,
		t:    t,
		g:    g,
		b:    b,
		y:    y,
		n:    n,
		c:    c,
		s:    s,
		app:  app,
		auth: auth,
	}
	return res
}
