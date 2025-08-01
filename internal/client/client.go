package client

import "github.com/google/wire"

// ProviderSet is client providers.
var ProviderSet = wire.NewSet(NewMinerClient, NewPorter, NewPorterClient, NewInprocPorter)
