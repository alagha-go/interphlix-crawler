package movies

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/gocolly/colly"
)


func CollectPages(PagesLength int) {
	for index:=1; index<PagesLength+1; index++ {
		CollectPageMovies(index)
		if index == 1 {
			if !MoviesAvailable() {
				break
			}
		}
	}
}


func CollectPageMovies(page int) {
	url := "https://tinyzonetv.to/movie?page=" + string(rune(page))
	collector := colly.NewCollector()

	collector.OnHTML(".film_list-wrap", CollectMovies)

	collector.Visit(url)
}


func CollectMovies(element *colly.HTMLElement) {
	element.ForEach(".flw-item", func(_ int, element *colly.HTMLElement) {
		var Movie Movie
        Movie.Title = element.ChildAttr("a", "title")
        Movie.ImageUrl = element.ChildAttr("img", "data-src")
        Movie.PageUrl = "https://tinyzonetv.to" + element.ChildAttr("a", "href")
		index := strings.Index(Movie.PageUrl, "free-")
    	Movie.Code = Movie.PageUrl[index+5:]
		PagesMovies = append(PagesMovies, Movie)
	})
}

func LoadDBPages() {
	data, err := ioutil.ReadFile(".DB/Movies/pages.json")
	HanleError(err)
	json.Unmarshal(data, &DBPages)
}

func MoviesAvailable() bool {
	for _, PageMovie := range PagesMovies {
		available := false
		for _, DBPageMovie := range DBPages {
			if PageMovie.Code == DBPageMovie.Code{
				available = true
			}
		}
		if !available {
			return available
		}
	}
	return true
}