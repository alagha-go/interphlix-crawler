package movies

import (
	"crawler/lib/types"
	"io/ioutil"
	"strings"

	"github.com/gocolly/colly"
)


var (
	Movies []types.Movie
	Position int
	Available int
)

func CollectAllPages(pages int) {
	for index:=1; index<pages+1; index++ {
		CollectPage(index)
	}
	SaveMovies()
	for index := range SavedMovies {
		Position = index
		CollectMovie(&SavedMovies[index])
		SetServers(&SavedMovies[index])
		SetID(&SavedMovies[index])
		AddServer(&SavedMovies[index])
		UploadMovie(&SavedMovies[index])
		SaveMovies()
	}
}

func CollectPage(number int) {
	url := "https://tinyzonetv.to/movie?page=" + string(rune(number))
	collector := colly.NewCollector()

	collector.OnHTML(".film_list-wrap", CollectMovies)

	collector.Visit(url)
}

func CollectMovies(element *colly.HTMLElement) {
	element.ForEach(".flw-item", func(_ int, element *colly.HTMLElement) {
		var Movie types.Movie
        Movie.Title = element.ChildAttr("a", "title")
        Movie.ImageUrl = element.ChildAttr("img", "data-src")
        Movie.PageUrl = "https://tinyzonetv.to" + element.ChildAttr("a", "href")
        Movie.PageUrl = strings.ReplaceAll(Movie.PageUrl, "/movie/", "/watch-movie/")
		index := strings.Index(Movie.PageUrl, "free-")
    	Movie.Code = Movie.PageUrl[index+5:]
		if !MovieExist(&Movie) {
			Movies = append(Movies, Movie)
		}
	})
}

func SaveMovies() {
	LoadData()
	SavedMovies = append(SavedMovies, Movies...)
	types.PrintGreen(len(SavedMovies))
	types.PrintPurple("Done Collecting All pages")
	data := types.JsonMarshal(SavedMovies)
	ioutil.WriteFile("./DB/movies.json", data, 0755)
	Movies = []types.Movie{}
}