package steam

import (
	"github.com/tuihub/librarian/internal/conf"

	"github.com/gocolly/colly/v2"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewSteam, NewStoreAPI, NewWebAPI)

type Steam struct {
	*StoreAPI
	*WebAPI
}

func NewSteam(c *colly.Collector, config *conf.Porter_Data) (*Steam, error) {
	if c == nil || config == nil || config.GetSteam() == nil {
		return new(Steam), nil
	}
	s, err := NewStoreAPI(c)
	if err != nil {
		return nil, err
	}
	w, err := NewWebAPI(c, config.GetSteam())
	if err != nil {
		return nil, err
	}
	return &Steam{
		StoreAPI: s,
		WebAPI:   w,
	}, nil
}

func (s *Steam) FeatureEnabled() bool {
	return s.StoreAPI != nil && s.WebAPI != nil
}
