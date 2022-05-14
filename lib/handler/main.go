package handler

import (
	"crawler/lib/movies"
	"crawler/lib/tvshows"
	"encoding/json"
	"log"
	"net/http"
)


func Main() {
	go movies.Main()
	go tvshows.Main()
	http.HandleFunc("/", GetStats)
}

func GetStats(res http.ResponseWriter, req *http.Request) {
	MovieType := Stats{"Movie", movies.LoopNumber, movies.PagesPosition, movies.MoviesPosition, movies.Available, movies.MoviesPosition-movies.Available, movies.UploadedMovies, movies.MoviesPosition-movies.UploadedMovies}
	TvShowType := Stats{"Tv-Show", tvshows.LoopNumber, tvshows.PagesPosition, tvshows.MoviesPosition, tvshows.Available, tvshows.MoviesPosition-tvshows.Available, tvshows.UploadedMovies, tvshows.MoviesPosition-tvshows.UploadedMovies}
	var StatsData = []Stats{MovieType, TvShowType}
	res.WriteHeader(200)
	json.NewEncoder(res).Encode(StatsData)
}

func HandleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}