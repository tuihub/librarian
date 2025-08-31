package model

// Subject represents a Bangumi subject (anime/manga/book/game/music)
type Subject struct {
	ID       int    `json:"id"`
	URL      string `json:"url"`
	Type     int    `json:"type"` // 1=book, 2=anime, 3=music, 4=game, 6=real
	Name     string `json:"name"`
	NameCN   string `json:"name_cn"`
	Summary  string `json:"summary"`
	Images   Images `json:"images"`
	InfoBox  []Info `json:"infobox"`
	EPS      int    `json:"eps"`
	Volumes  int    `json:"volumes"`
	Locked   bool   `json:"locked"`
	NSFW     bool   `json:"nsfw"`
	Date     string `json:"date"`
	Platform string `json:"platform"`
	Tags     []Tag  `json:"tags"`
	Rating   Rating `json:"rating"`
}

type Images struct {
	Small  string `json:"small"`
	Grid   string `json:"grid"`
	Large  string `json:"large"`
	Medium string `json:"medium"`
	Common string `json:"common"`
}

type Info struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Tag struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type Rating struct {
	Rank  int     `json:"rank"`
	Total int     `json:"total"`
	Count map[string]int `json:"count"`
	Score float64 `json:"score"`
}

// SearchSubjectsResponse for search API
type SearchSubjectsResponse struct {
	Results int       `json:"results"`
	List    []Subject `json:"list"`
}

// SubjectType represents the type of a Bangumi subject
type SubjectType int

const (
	SubjectTypeBook  SubjectType = 1
	SubjectTypeAnime SubjectType = 2
	SubjectTypeMusic SubjectType = 3
	SubjectTypeGame  SubjectType = 4
	SubjectTypeReal  SubjectType = 6
)