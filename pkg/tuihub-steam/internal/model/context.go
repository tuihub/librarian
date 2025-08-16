package model

type PorterContext struct {
	APIKey string `json:"api_key" jsonschema:"title=Steam API Key,description=Your Steam Web API key"`
}

type GetAccountConfig struct {
	AccountID string `json:"account_id" jsonschema:"title=Account ID"`
}

type GetAppInfoConfig struct {
	AppID string `json:"app_id" jsonschema:"title=App ID"`
}

type SearchAppInfoConfig struct {
	NameLike string `json:"name_like" jsonschema:"title=Name Like"`
}
