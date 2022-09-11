package inprocgrpc

import (
	"github.com/fullstorydev/grpchan/inprocgrpc"
	"github.com/google/wire"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"
)

var ProviderSet = wire.NewSet(NewInprocMapperChannel, NewInprocSearcherChannel, NewInprocPorterChannel)

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
