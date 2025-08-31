package model

type PorterContext struct {
	Token string `json:"token" jsonschema:"title=Bangumi API Token,description=Your Bangumi API access token,required"`
}

type GetAppInfoConfig struct {
	AppID string `json:"app_id" jsonschema:"title=Subject ID,description=Bangumi subject ID"`
}

type SearchAppInfoConfig struct {
	NameLike string `json:"name_like" jsonschema:"title=Name Like,description=Search query for anime/manga titles"`
}