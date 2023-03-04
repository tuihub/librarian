package service

import (
	"context"
	"os"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizangela"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizbinah"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizyesod"
	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
)

type LibrarianSephirahServiceService struct {
	pb.UnimplementedLibrarianSephirahServiceServer

	converter converter.Converter

	t *biztiphereth.Tiphereth
	g *bizgebura.Gebura
	b *bizbinah.Binah
	y *bizyesod.Yesod
}

func NewLibrarianSephirahServiceService(
	_ *bizangela.Angela,
	t *biztiphereth.Tiphereth,
	g *bizgebura.Gebura,
	b *bizbinah.Binah,
	y *bizyesod.Yesod,
) pb.LibrarianSephirahServiceServer {
	if _, exist := os.LookupEnv("CREATE_ADMIN"); exist {
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
		converter: converter.NewConverter(),
		t:         t,
		g:         g,
		b:         b,
		y:         y,
	}
}
