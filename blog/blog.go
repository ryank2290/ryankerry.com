package main

import (
	"net/http"

	"github.com/ryank90/utilities/blog"
)

const hostname = "ryankerry.com" // Default hostname for the blog server.

var config = blog.Config{
	Hostname:     hostname,
	BaseURL:      "//" + hostname,
	HomeArticles: 1, // Number of articles to display on the homepage.
	FeedArticles: 1, // Number of articles to display on the ATOM and JSON feeds.
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
