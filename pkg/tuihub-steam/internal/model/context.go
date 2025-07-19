package model

type PorterContext struct {
	APIKey string `json:"api_key" jsonschema:"title=Steam API Key,description=Your Steam Web API key"`
}
