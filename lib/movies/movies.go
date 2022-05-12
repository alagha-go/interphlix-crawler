package movies


func CollectAllMovies() {
	for index := range PagesMovies {
		CollectMovie(index)
	}
}

func CollectMovie(index int) {
	Movie := PagesMovies[index]
	if !Movie.Collected {
		Movie.CollectMovieContent()
	}
}