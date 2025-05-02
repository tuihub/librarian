package internal

type PorterContext struct {
	Token string `json:"token" jsonschema:"title=Token"`
}

type PushFeedItems struct {
	ChannelID int64 `json:"channel_id,string" jsonschema:"title=Channel ID"`
}
