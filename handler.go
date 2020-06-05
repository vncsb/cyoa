package cyoa

import (
	"net/http"
	"strings"
	"text/template"
)

const defaultIntro = "intro"

type StoryHandler struct {
	story         Story
	storyTemplate template.Template
}

func NewStoryHandler(story Story, tmpl template.Template) *StoryHandler {
	sh := StoryHandler{}
	sh.story = story
	sh.storyTemplate = tmpl
	if sh.story.IntroTitle == "" {
		sh.story.IntroTitle = defaultIntro
	}

	return &sh
}

func (sh *StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	chapter, ok := sh.story.Chapters[path]
	if !ok {
		chapter = sh.story.Chapters[sh.story.IntroTitle]
	}

	err := sh.storyTemplate.Execute(w, chapter)

	if err != nil {
		http.Error(w, "Something went wrong...", http.StatusInternalServerError)
	}
}
