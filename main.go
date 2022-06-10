package main

import (
	"crawler/lib/crawler/movies"
	// "crawler/lib/handler"
	// "net/http"
)

var (
	PORT = ":7000"
)

func main() {
	movies.LoadMovies()
	// handler.Main()
	// http.ListenAndServe(PORT, nil)
}