package main

//import (
//	"net/http"
//
//	"github.com/ryank90/utilities/blog"
//)
//
//func init() {
//	config.ContentPath = "articles/"
//	config.ThemePath = "theme/"
//
//	s, err := blog.NewServer(config)
//	if err != nil {
//		panic(err)
//	}
//
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Strict-Transport-Security", "max-age=31536000; preload")
//		s.ServeHTTP(w, r)
//	})
//}
