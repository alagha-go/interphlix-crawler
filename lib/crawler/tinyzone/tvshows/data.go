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

func UploadTvShow(TvShow *types.Movie) {
	var newTvShow types.Movie
	data, _, _ := PostRequest("https://s1.interphlix.com/movies/upload", types.JsonMarshal(TvShow), false)
	err := json.Unmarshal(data, &newTvShow)
	if err != nil {
		return
	}
	TvShow.Uploaded = true
}