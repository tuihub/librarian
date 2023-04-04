package service

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizangela"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizbinah"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biznetzach"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizyesod"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/internal/lib/libapp"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
)

type LibrarianSephirahServiceService struct {
	pb.UnimplementedLibrarianSephirahServiceServer

	t        *biztiphereth.Tiphereth
	g        *bizgebura.Gebura
	b        *bizbinah.Binah
	y        *bizyesod.Yesod
	n        *biznetzach.Netzach
	authFunc func(context.Context) (context.Context, error)
}

func NewLibrarianSephirahServiceService(
	_ *bizangela.Angela,
	t *biztiphereth.Tiphereth,
	g *bizgebura.Gebura,
	b *bizbinah.Binah,
	y *bizyesod.Yesod,
	n *biznetzach.Netzach,
	app *libapp.Settings,
	authFunc func(context.Context) (context.Context, error),
) pb.LibrarianSephirahServiceServer {
	if enable, err := app.GetEnvBool(libapp.EnvCreateAdmin); err == nil && enable {
		t.CreateDefaultAdmin(context.Background(), &modeltiphereth.User{
			ID:       0,
			UserName: "admin",
			PassWord: "admin",
			Type:     0,
			Status:   0,
		})
	}
	return &LibrarianSephirahServiceService{
		UnimplementedLibrarianSephirahServiceServer: pb.UnimplementedLibrarianSephirahServiceServer{},
		t:        t,
		g:        g,
		b:        b,
		y:        y,
		n:        n,
		authFunc: authFunc,
	}
}
