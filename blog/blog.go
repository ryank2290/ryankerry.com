package main

import (
	"net/http"

	"golang.org/x/tools/godoc/static"

	"strings"
	"time"

	"github.com/ryank90/utilities/blog"
)

const hostname = "ryankerry.com" // Default hostname for the blog server.

var config = blog.Config{
	Hostname:     hostname,
	BaseURL:      "//" + hostname,
	HomeArticles: 10, // Number of content to display on the homepage.
	FeedArticles: 10, // Number of content to display on the ATOM and JSON feeds.
	FeedTitle:    "Ryan Kerry's Blog",
}

func init() {
	redirect := func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	http.HandleFunc("/blog", redirect)
	http.HandleFunc("/blog/", redirect)

	http.Handle("/fonts/", http.FileServer(http.Dir("static")))
	http.Handle("/styles/", http.FileServer(http.Dir("static")))
	http.Handle("/scripts/", http.FileServer(http.Dir("static")))
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path
	b, ok := static.Files[name]
	if !ok {
		http.NotFound(w, r)
		return
	}
	http.ServeContent(w, r, name, time.Time{}, strings.NewReader(b))
}
