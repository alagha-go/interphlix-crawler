package handler

import (
	"crawler/lib/crawler/tvshows"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetAllTvShows(res http.ResponseWriter, req *http.Request) {
	data, err := ioutil.ReadFile("./DB/Tvshows/tvshows.json")
	HandleError(err)
	res.WriteHeader(200)
	res.Write(data)
}

func GetUnAvailableTvShows(res http.ResponseWriter, req *http.Request) {
	var TvShows []tvshows.TvShow
	var UnAvailableTvShows []tvshows.TvShow
	data, err := ioutil.ReadFile("./DB/Tvshows/tvshows.json")
	HandleError(err)
	json.Unmarshal(data, &TvShows)
	for _, TvShow := range TvShows {
		if !TvShow.Available && TvShow.Collected {
			UnAvailableTvShows = append(UnAvailableTvShows, TvShow)
		}
	}
	json.NewEncoder(res).Encode(UnAvailableTvShows)
}


func GetUnUploadedTvShows(res http.ResponseWriter, req *http.Request) {
	var TvShows []tvshows.TvShow
	var UnUploadedTvShows []tvshows.TvShow
	data, err := ioutil.ReadFile("./DB/Tvshows/tvshows.json")
	HandleError(err)
	json.Unmarshal(data, &TvShows)
	for _, TvShow := range TvShows {
		if !TvShow.Uploaded && TvShow.Collected {
			UnUploadedTvShows = append(UnUploadedTvShows, TvShow)
		}
	}
	json.NewEncoder(res).Encode(UnUploadedTvShows)
}