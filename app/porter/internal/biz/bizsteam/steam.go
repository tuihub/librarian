package bizsteam

import (
	"github.com/tuihub/librarian/app/porter/internal/client/steam"
	"github.com/tuihub/librarian/app/porter/internal/data"
)

type SteamUseCase struct {
	c *steam.Steam
}

func NewSteamUseCase(client *steam.Steam, _ *data.Data) *SteamUseCase {
	return &SteamUseCase{c: client}
}
