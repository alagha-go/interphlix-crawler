package tvshows

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func CollectPages(PagesLength int) {
	TotalNumberOfPages = PagesLength
	for index:=1; index < PagesLength+1; index++ {
		CurrentPageNumber = index
		CollectPageTvShows(index)
		if index % 20 == 0 {
			PrintYellow(fmt.Sprintf("TvShows   :%d",index))
		}
	}
	SaveTvShows()
	PrintBlue(len(TvShows))
	PrintCyan(EpisodesLength)
	PrintGreen("done collecting all the pages data")

	for index := range TvShows {
		CurrentMovie = index+1
		if !TvShows[index].Collected {
			TvShows[index].GetSeasons()
			TvShows[index].Collected = true
		}
		if !TvShows[index].Uploaded {
			TvShows[index].Upload()
		}else {
			TvShows[index].Update()
		}
		SaveTvShows()
	}
}


func CollectPageTvShows(page int) {
	url := "https://tinyzonetv.to/tv-show?page=" + strconv.Itoa(page)
	collector := colly.NewCollector()

	collector.OnHTML(".film_list-wrap", CollectTvShows)

	collector.Visit(url)
}


func CollectTvShows(element *colly.HTMLElement) {
	element.ForEach(".flw-item", func(pos int, element *colly.HTMLElement) {
		CurrentPageCollectedMovies = pos
		var TvShow TvShow
		TvShow.SetTvShowID()
		TvShow.ID = primitive.NewObjectID()
        TvShow.Title = element.ChildAttr("a", "title")
        TvShow.ImageUrl = element.ChildAttr("img", "data-src")
        TvShow.PageUrl = "https://tinyzonetv.to" + element.ChildAttr("a", "href")
		index := strings.Index(TvShow.PageUrl, "free-")
    	TvShow.Code = TvShow.PageUrl[index+5:]
		TvShow.Type = "Tv-Show"
		if !TvShow.Exists() {
			TvShow.CollectTvShowContent()
			TvShows = append(TvShows, TvShow)
		}
	})
}


func SaveTvShows() {
	data := JsonMarshal(TvShows)
	ioutil.WriteFile("./DB/Tvshows/tvshows.json", data, 0755)
}