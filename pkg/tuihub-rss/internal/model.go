package internal

type PullRSSConfig struct {
	URL string `json:"url" jsonschema:"title=URL"`
}

type ServeRSSConfig struct {
	Title string `json:"title" jsonschema:"title=Title"`
}
