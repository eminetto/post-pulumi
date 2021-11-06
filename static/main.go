package main

import (
	"log"
	"net/http"
)

func main() {
	r := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./"))
	r.Handle("/", http.StripPrefix("/", fileServer))
	s := &http.Server{
		Addr:    ":80",
		Handler: r,
	}
	log.Fatal(s.ListenAndServe())
}
