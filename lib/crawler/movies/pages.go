package movies

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)


func CollectPages(PagesLength int) {
	TotalNumberOfPages = PagesLength
	for index:=1; index<PagesLength+1; index++ {
		CurrentPageNumber = index
		CollectPageMovies(index)
		if index % 50 == 0 {
			PrintYellow(fmt.Sprintf("Movies    :%d",index))
		}
	}
	PrintBlue(len(Movies))
	SaveMovies()
	PrintGreen("done collecting all the pages data")

	for index := range Movies {
		CurrentMovie = index+1
		if !Movies[index].Collected {
			Movies[index].SetServers()
			Movies[index].Collected = true
		}
		if !Movies[index].Uploaded {
			Movies[index].Upload()
		}
		SaveMovies()
	}
}


func CollectPageMovies(page int) {
	url := "https://tinyzonetv.to/movie?page=" + strconv.Itoa(page)
	collector := colly.NewCollector()

	collector.OnHTML(".film_list-wrap", CollectMovies)

	collector.Visit(url)
}


func CollectMovies(element *colly.HTMLElement) {
	element.ForEach(".flw-item", func(pos int, element *colly.HTMLElement) {
		CurrentPageCollectedMovies = pos
		var Movie Movie
		Movie.SetMovieID()
        Movie.Title = element.ChildAttr("a", "title")
        Movie.ImageUrl = element.ChildAttr("img", "data-src")
        Movie.PageUrl = "https://tinyzonetv.to" + element.ChildAttr("a", "href")
		index := strings.Index(Movie.PageUrl, "free-")
    	Movie.Code = Movie.PageUrl[index+5:]
		if !Movie.Exists() {
			Movie.CollectMovieContent()
			Movies = append(Movies, Movie)
		}
	})
}


func SaveMovies() {
	data := JsonMarshal(Movies)
	ioutil.WriteFile("./DB/Movies/movies.json", data, 0755)
}