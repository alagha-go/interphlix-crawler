package tvshows


import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/cheggaaa/pb/v3"
	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func CollectPages(PagesLength int) {
	bar := pb.StartNew(PagesLength-1)
	for index:=1; index < PagesLength+1; index++ {
		CollectPageMovies(index)
		bar.Increment()
		PagesPosition = index
	}
	bar.Finish()
	SavePagesData()
	PrintGreen("done collecting all the pages data")
}


func CollectPageMovies(page int) {
	url := "https://tinyzonetv.to/tv-show?page=" + strconv.Itoa(page)
	collector := colly.NewCollector()

	collector.OnHTML(".film_list-wrap", CollectMovies)

	collector.Visit(url)
}


func CollectMovies(element *colly.HTMLElement) {
	element.ForEach(".flw-item", func(_ int, element *colly.HTMLElement) {
		var Movie Movie
		Movie.ID = primitive.NewObjectID()
        Movie.Title = element.ChildAttr("a", "title")
        Movie.ImageUrl = element.ChildAttr("img", "data-src")
        Movie.PageUrl = "https://tinyzonetv.to" + element.ChildAttr("a", "href")
		index := strings.Index(Movie.PageUrl, "free-")
    	Movie.Code = Movie.PageUrl[index+5:]
		PagesTvShows = append(PagesTvShows, Movie)
	})
}

func LoadDBPages() {
	data, err := ioutil.ReadFile("./DB/Tvshows/pages.json")
	HanleError(err)
	json.Unmarshal(data, &DBPages)
}



func SavePagesData() {
	data := JsonMarshal(PagesTvShows)
	ioutil.WriteFile("./DB/Tvshows/pages.json", data, 0755)
}