package movies

import (
	"encoding/json"
	"strings"

	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


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


func (Movie *Movie) SetMovieID() {
    ID := primitive.NewObjectID()
    for _, movie := range Movies {
        if ID == movie.ID {
            Movie.SetMovieID()
        }
    }
    Movie.ID = ID
}


func (Movie *Movie) Exists() bool {
	for index := range Movies {
		if Movie.Code == Movies[index].Code {
			return true
		}
	}
	return false
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
