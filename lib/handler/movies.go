package handler

import (
	"io/ioutil"
	"net/http"
)

func GetAllMovies(res http.ResponseWriter, req *http.Request) {
	data, err := ioutil.ReadFile("./DB/Movies/movies.json")
	HandleError(err)
	res.WriteHeader(200)
	res.Write(data)
}