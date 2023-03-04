package modelangela

import "github.com/tuihub/librarian/internal/model"

type PullSteamAccountAppRelation struct {
	ID      model.InternalID
	SteamID string
}

type PullSteamApp struct {
	ID    model.InternalID
	AppID string
}
