package handler

import (
	"crawler/lib/crawler/movies"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetAllMovies(res http.ResponseWriter, req *http.Request) {
	data, err := ioutil.ReadFile("./DB/Movies/movies.json")
	HandleError(err)
	res.WriteHeader(200)
	res.Write(data)
}

func GetUnAvailableMovies(res http.ResponseWriter, req *http.Request) {
	var Movies []movies.Movie
	var UnAvailableMovies []movies.Movie
	data, err := ioutil.ReadFile("./DB/Movies/movies.json")
	HandleError(err)
	json.Unmarshal(data, &Movies)
	for _, Movie := range Movies {
		if !Movie.Available {
			UnAvailableMovies = append(UnAvailableMovies, Movie)
		}
	}
	json.NewEncoder(res).Encode(UnAvailableMovies)
}

func GetUnUploadedMovies(res http.ResponseWriter, req *http.Request) {
	var Movies []movies.Movie
	var UnUploadeMovies []movies.Movie
	data, err := ioutil.ReadFile("./DB/Movies/movies.json")
	HandleError(err)
	json.Unmarshal(data, &Movies)
	for _, Movie := range Movies {
		if !Movie.Uploaded {
			UnUploadeMovies = append(UnUploadeMovies, Movie)
		}
	}
	json.NewEncoder(res).Encode(UnUploadeMovies)
}