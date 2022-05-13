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
		Movie.SetMovieID()
		SavePagesData()
		Movie.Collected = true
		Movie.IsAvailable()
		Movie.Upload()
		Movies = append(Movies, Movie)
		SaveMovies()
		MoviesPosition++
		if Movie.Available {
			Available++
		}
		if Movie.Uploaded {
			UploadedMovies++
		}
	}else {
		Movie.SetDBMovie()
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

func (Movie *Movie) IsAvailable() {
	if Movie.Server.Name == "Streamlare" {
		Movie.Available = true
	}
}


func (Movie *Movie) SetDBMovie() {
	for index := range DBMovies {
		if Movie.ID == DBMovies[index].ID {
			Movie = &DBMovies[index]
		}
	}
}


func LoadDBMovies() {
	data, err := ioutil.ReadFile("./DB/Movies/movies.json")
	HanleError(err)
	json.Unmarshal(data, &DBMovies)
}