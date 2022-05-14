package handler

import (
	"io/ioutil"
	"net/http"
)

func GetAllTvShows(res http.ResponseWriter, req *http.Request) {
	data, err := ioutil.ReadFile("./DB/Tvshows/tvshows.json")
	HandleError(err)
	res.WriteHeader(200)
	res.Write(data)
}