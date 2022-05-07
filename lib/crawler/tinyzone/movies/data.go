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

func UploadMovie(Movie *types.Movie) {
	var newMovie types.Movie
	data, _, _ := PostRequest("https://s1.interphlix.com/movies/upload", types.JsonMarshal(Movie), false)
	err := json.Unmarshal(data, &newMovie)
	if err != nil {
		return
	}
	Movie.Uploaded = true
}