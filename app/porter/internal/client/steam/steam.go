package steam

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
