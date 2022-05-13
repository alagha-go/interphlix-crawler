package tvshows

import (
	"encoding/json"
	"io/ioutil"
)


func UploadUnUploadedTvShows() {
	var TvShows []Movie
	data, err := ioutil.ReadFile("./DB/TvShows/tvshows.json")
	HanleError(err)
	json.Unmarshal(data, &TvShows)
	for index := range TvShows {
		if !TvShows[index].Uploaded && TvShows[index].Collected{
			TvShows[index].Upload()
			if !TvShows[index].Uploaded {
				TvShows[index].SetMovieID()
				TvShows[index].Upload()
			}
			ioutil.WriteFile("./DB/Tvshows/tvshows.json", JsonMarshal(TvShows), 0755)
		}
	}
}