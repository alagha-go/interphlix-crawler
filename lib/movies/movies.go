package movies

import (
	"encoding/json"
	"io/ioutil"
)


func CollectAllMovies() {
	for index := range PagesMovies {
		CollectMovie(index)
	}
}

func CollectMovie(index int) {
	Movie := PagesMovies[index]
	if !Movie.Collected {
		Movie.CollectMovieContent()
		PagesMovies[index].Collected = true
		SavePagesData()
		Movie.Upload()
		Movies = append(Movies, Movie)
		SaveMovies()
	}
}


func SaveMovies() {
	data := JsonMarshal(Movies)
	ioutil.WriteFile("./DB/Movies/movies.json", data, 0755)
}


func (movie *Movie) Upload() {
	var newMovie Movie
	data, _, _ := PostRequest("https://s1.interphlix.com/movies/upload", JsonMarshal(movie), false)
	err := json.Unmarshal(data, &newMovie)
	if err != nil {
		return
	}
	movie.Uploaded = true
}