package handler

import (
	"crawler/lib/crawler/movies"
	"crawler/lib/crawler/tvshows"
	"encoding/json"
	"log"
	"net/http"
)


func Main() {
	go movies.Main()
	go tvshows.Main()
	http.HandleFunc("/", GetStats)
	http.HandleFunc("/movies/pages", MoviePages)
	http.HandleFunc("/tv-shows/pages", TvShowsPages)
	http.HandleFunc("/movies/all", GetAllMovies)
	http.HandleFunc("/tv-shows/all", GetAllTvShows)
	http.HandleFunc("/movies/unavailable", GetUnAvailableMovies)
	http.HandleFunc("/tv-shows/unavailable", GetUnAvailableTvShows)
	http.HandleFunc("/movies/unuploaded", GetUnUploadedMovies)
	http.HandleFunc("/tv-shows/unuploaded", GetUnUploadedTvShows)
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