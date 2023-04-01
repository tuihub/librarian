package inprocgrpc

import (
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"

	"github.com/fullstorydev/grpchan/inprocgrpc"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewInprocClients)

type InprocClients struct {
	Mapper   mapper.LibrarianMapperServiceClient
	Searcher searcher.LibrarianSearcherServiceClient
	Porter   porter.LibrarianPorterServiceClient
}

func NewInprocClients(
	m mapper.LibrarianMapperServiceServer,
	s searcher.LibrarianSearcherServiceServer,
	p porter.LibrarianPorterServiceServer,
) *InprocClients {
	return &InprocClients{
		Mapper:   NewInprocMapperChannel(m),
		Searcher: NewInprocSearcherChannel(s),
		Porter:   NewInprocPorterChannel(p),
	}
}

func NewInprocMapperChannel(s mapper.LibrarianMapperServiceServer) mapper.LibrarianMapperServiceClient {
	channel := inprocgrpc.Channel{}
	mapper.RegisterLibrarianMapperServiceServer(&channel, s)
	cli := mapper.NewLibrarianMapperServiceClient(&channel)
	return cli
}

func NewInprocSearcherChannel(s searcher.LibrarianSearcherServiceServer) searcher.LibrarianSearcherServiceClient {
	channel := inprocgrpc.Channel{}
	searcher.RegisterLibrarianSearcherServiceServer(&channel, s)
	cli := searcher.NewLibrarianSearcherServiceClient(&channel)
	return cli
}

func NewInprocPorterChannel(s porter.LibrarianPorterServiceServer) porter.LibrarianPorterServiceClient {
	channel := inprocgrpc.Channel{}
	porter.RegisterLibrarianPorterServiceServer(&channel, s)
	cli := porter.NewLibrarianPorterServiceClient(&channel)
	return cli
}
