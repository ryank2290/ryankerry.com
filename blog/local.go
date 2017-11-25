package main

import (
	"flag"
	"net/http"

	"log"

	"net"

	"github.com/ryank90/utilities/blog"
)

var (
	httpAddr     = flag.String("http", "localhost:8080", "HTTP listen address")
	contentPath  = flag.String("content", "content/", "Path to article files")
	templatePath = flag.String("template", "template/", "Path to template files")
	staticPath   = flag.String("static", "static/", "Path to static files")
	reload       = flag.Bool("reload", false, "Reload content on each page load")
)

func main() {
	flag.Parse()

	config.ContentPath = *contentPath
	config.TemplatePath = *templatePath

	mux, err := newServer(*reload, *staticPath, config)

	if err != nil {
		log.Fatal(err)
	}

	ln, err := net.Listen("tcp", *httpAddr)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Listening on addr", *httpAddr)
	log.Fatal(http.Serve(ln, mux))
}

func newServer(reload bool, staticPath string, config blog.Config) (http.Handler, error) {
	mux := http.NewServeMux()

	if reload {
		mux.HandleFunc("/", reloadServer)
	} else {
		s, err := blog.NewServer(config)
		if err != nil {
			return nil, err
		}
		mux.Handle("/", s)
	}

	fs := http.FileServer(http.Dir(staticPath))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	return mux, nil
}

func reloadServer(w http.ResponseWriter, r *http.Request) {
	s, err := blog.NewServer(config)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	s.ServeHTTP(w, r)
}
