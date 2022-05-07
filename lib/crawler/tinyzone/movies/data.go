package movies

import (
	"crawler/lib/types"
	"encoding/json"
	"io/ioutil"
)


func LoadData() {
	data, err := ioutil.ReadFile("./DB/movies.json")
	HanleError(err)
	json.Unmarshal(data, &Movies)
}


func MovieExist(Movie *types.Movie) bool {
	for _, movie := range Movies {
		if movie.Code == Movie.Code {
			return true
		}
	}
	return false
}