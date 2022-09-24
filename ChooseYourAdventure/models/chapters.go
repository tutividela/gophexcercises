package models

type Story map[string]Chapter

type Chapter struct {
	Title   string       `json:"title"`
	Story   []string     `json:"story"`
	Options []OptionItem `json:"options"`
}

type OptionItem struct {
	Text         string `json:"text"`
	ChapterTitle string `json:"arc"`
}