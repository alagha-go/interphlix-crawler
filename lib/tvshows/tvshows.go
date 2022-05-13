package tvshows

import (
	"encoding/json"
	"io/ioutil"
)

func CollectAllTvShows() {
	for index := range PagesTvShows {
		CollectTvShow(index)
	}
}

func CollectTvShow(index int) {
	TvShow := PagesTvShows[index]
	if !TvShow.Collected {
		TvShow.CollectTvShowContent()
		PagesTvShows[index].Collected = true
		SavePagesData()
		TvShow.Collected = true
		TvShow.IsAvailable()
		TvShow.Upload()
		TvShows = append(TvShows, TvShow)
		SaveTvShows()
		MoviesPosition++
		if TvShow.Available {
			Available++
		}
		if TvShow.Uploaded {
			UploadedMovies++
		}
	}
}


func SaveTvShows() {
	data := JsonMarshal(TvShows)
	ioutil.WriteFile("./DB/Tvshows/tvshows.json", data, 0755)
}


func (movie *Movie) Upload() {
	var newTvShow Movie
	data, _, _ := PostRequest("https://s1.interphlix.com/movies/upload", JsonMarshal(movie), false)
	err := json.Unmarshal(data, &newTvShow)
	if err != nil {
		return
	}
	movie.Uploaded = true
}


func (TvShow *Movie) IsAvailable() {
	for _, Season := range TvShow.Seasons {
		for _, Episode := range Season.Episodes {
			if Episode.Server.Name == "Streamlare" {
				TvShow.Available = true
				return
			}
		}
	}
}