package steam

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewSteam, NewStoreAPI, NewWebAPI)

type Steam struct {
	StoreAPI *StoreAPI
	WebAPI   *WebAPI
}

func NewSteam(s *StoreAPI, w *WebAPI) *Steam {
	return &Steam{
		StoreAPI: s,
		WebAPI:   w,
	}
}
