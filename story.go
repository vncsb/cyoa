package cyoa

import (
	"encoding/json"
	"io"
)

func JsonStory(r io.Reader, intro string) (Story, error) {
	d := json.NewDecoder(r)
	var chapters map[string]Chapter
	if err := d.Decode(&chapters); err != nil {
		return Story{}, err
	}

	story := Story{
		Chapters:   chapters,
		IntroTitle: intro,
	}

	return story, nil
}

type Story struct {
	Chapters   map[string]Chapter
	IntroTitle string
}
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
