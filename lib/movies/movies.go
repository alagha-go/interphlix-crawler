package movies

import (
	"strings"

	"github.com/gocolly/colly"
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
	}
}


func (Movie *Movie)CollectMovieContent() {
    collector := colly.NewCollector()

    collector.OnHTML(".description", func(element *colly.HTMLElement){
        Movie.Description = element.Text
        Movie.Description = strings.ReplaceAll(Movie.Description, "\n", "")
        Movie.Description = strings.ReplaceAll(Movie.Description, "  ", "")
        Movie.Description = strings.TrimPrefix(Movie.Description, " ")
        Movie.Description = strings.TrimSuffix(Movie.Description, " ")
    })

    
    collector.OnHTML(".elements", Movie.SetElements)
    
    collector.Visit(Movie.PageUrl)
}


func (Movie *Movie)SetElements(element *colly.HTMLElement) {
    functions := []func(element *colly.HTMLElement){}
    functions = append(functions,  Movie.SetReleased)
    functions = append(functions,  Movie.SetGenre)
    functions = append(functions,  Movie.SetCasts)
    functions = append(functions,  Movie.SetDuration)
    functions = append(functions,  Movie.SetCountries)
    functions = append(functions,  Movie.SetProducers)
    element.ForEach(".row-line", func(index int, element *colly.HTMLElement){
        functions[index](element)
    })
}