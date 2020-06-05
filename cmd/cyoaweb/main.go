package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/vncsb/cyoa"
)

func main() {
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Println(*filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f, "intro")
	if err != nil {
		panic(err)
	}

	tmpl, err := template.ParseFiles(`C:\Users\vinic\Documents\Go\choose-your-own-adventure\tmpl\chapter.html`)

	sh := cyoa.NewStoryHandler(story, *tmpl)

	http.ListenAndServe(":8080", sh)
}
