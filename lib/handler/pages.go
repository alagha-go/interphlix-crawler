package handler

import (
	"io/ioutil"
	"net/http"
)

func MoviePages(res http.ResponseWriter, req *http.Request) {
	data, err := ioutil.ReadFile("./DB/Movies/pages.json")
	HandleError(err)
	res.WriteHeader(200)
	res.Write(data)
}

func TvShowsPages(res http.ResponseWriter, req *http.Request) {
	data, err := ioutil.ReadFile("./DB/Tvshows/pages.json")
	HandleError(err)
	res.WriteHeader(200)
	res.Write(data)
}