package inprocgrpc

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewInprocClients)

type InprocClients struct {
}

func NewInprocClients() *InprocClients {
	return &InprocClients{}
}
