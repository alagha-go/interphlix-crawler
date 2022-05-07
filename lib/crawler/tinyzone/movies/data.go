package movies

import (
	"crawler/lib/types"
	"encoding/json"
	"io/ioutil"
)

var (
	SavedMovies []types.Movie
)

func LoadData() {
	data, err := ioutil.ReadFile("./DB/movies.json")
	HanleError(err)
	json.Unmarshal(data, &SavedMovies)
}


func MovieExist(Movie *types.Movie) bool {
	for _, movie := range SavedMovies {
		if movie.Code == Movie.Code {
			return true
		}
	}
	return false
}