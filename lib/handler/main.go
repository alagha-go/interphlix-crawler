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
	http.HandleFunc("/movies/all", GetAllMovies)
	http.HandleFunc("/tv-shows/all", GetAllTvShows)
	http.HandleFunc("/movies/unavailable", GetUnAvailableMovies)
	http.HandleFunc("/tv-shows/unavailable", GetUnAvailableTvShows)
	http.HandleFunc("/movies/unuploaded", GetUnUploadedMovies)
	http.HandleFunc("/tv-shows/unuploaded", GetUnUploadedTvShows)
}

func GetStats(res http.ResponseWriter, req *http.Request) {
	Statistics := []Statistics{
		{
			Type: "Movie",
			LoopNumber: movies.LoopNumber,
			PagesLength: movies.TotalNumberOfPages,
			CurrentPageNumber: movies.CurrentPageNumber,
			CurrentPageCollectedMovies: movies.CurrentPageCollectedMovies,
			TotalNumberOfMovies: len(movies.Movies),
			CurrentMovie: movies.CurrentMovie,
			Available: movies.Available,
			UnAvailable: movies.CurrentMovie-movies.Available,
			Uploaded: movies.Uploaded,
			UnUploaded: movies.CurrentMovie-movies.Uploaded,
		},
		{
			Type: "Tv-Show",
			LoopNumber: tvshows.LoopNumber,
			PagesLength: tvshows.TotalNumberOfPages,
			CurrentPageNumber: tvshows.CurrentPageNumber,
			CurrentPageCollectedMovies: tvshows.CurrentPageCollectedMovies,
			TotalNumberOfMovies: len(tvshows.TvShows),
			CurrentMovie: tvshows.CurrentMovie,
			NumberOfEpisodes: tvshows.EpisodesLength,
			Available: tvshows.Available,
			UnAvailable: tvshows.CurrentMovie-tvshows.Available,
			Uploaded: tvshows.Uploaded,
			UnUploaded: tvshows.CurrentMovie-tvshows.Uploaded,
		},
	}

	json.NewEncoder(res).Encode(Statistics)
}

func HandleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}