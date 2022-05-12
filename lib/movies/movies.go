package movies

import "io/ioutil"


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
		Movies = append(Movies, Movie)
		SaveMovies()
	}
}


func SaveMovies() {
	data := JsonMarshal(Movies)
	ioutil.WriteFile("./DB/Movies/movies.json", data, 0755)
}