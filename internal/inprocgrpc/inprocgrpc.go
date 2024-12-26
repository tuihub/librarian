package inprocgrpc

import (
	miner "github.com/tuihub/protos/pkg/librarian/miner/v1"

	"github.com/fullstorydev/grpchan/inprocgrpc"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewInprocClients)

type InprocClients struct {
	Miner miner.LibrarianMinerServiceClient
}

func NewInprocClients(
	mi miner.LibrarianMinerServiceServer,
) *InprocClients {
	return &InprocClients{
		Miner: NewInprocMinerChannel(mi),
	}
}

func NewInprocMinerChannel(s miner.LibrarianMinerServiceServer) miner.LibrarianMinerServiceClient {
	channel := inprocgrpc.Channel{}
	miner.RegisterLibrarianMinerServiceServer(&channel, s)
	cli := miner.NewLibrarianMinerServiceClient(&channel)
	return cli
}
