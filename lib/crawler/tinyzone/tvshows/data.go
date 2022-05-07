package tvshows

import (
	"crawler/lib/types"
	"encoding/json"
	"io/ioutil"
)


var (
	SavedTvShows []types.Movie
)

func LoadData() {
	data, err := ioutil.ReadFile("./DB/tvshows.json")
	HanleError(err)
	json.Unmarshal(data, &SavedTvShows)
}


func TvShowExist(TvShow *types.Movie) bool {
	for _, movie := range SavedTvShows {
		if movie.Code == TvShow.Code {
			return true
		}
	}
	return false
}