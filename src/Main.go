package main

import (
	"flag"
	"log"
	"net/http"
	"time"
)

func main() {
	var dir string

	flag.StringVar(&dir, "dir", "../static", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	router := NewRouter()

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
