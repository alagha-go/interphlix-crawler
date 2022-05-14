package tvshows

import (
	"bytes"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

var (
	PagesTvShows []Movie
	DBPages		[]Movie
	TvShows		[]Movie
	DBTvShows	[]Movie
	LoopNumber int = 1
	PagesPosition int
	MoviesPosition int
	Available int
	UploadedMovies int
)

func Main() {
	for {
		LoadDBPages()
		LoadDBTvShows()
		CollectPages(GetNumberOfPages())
		CollectAllTvShows()
		UploadUnUploadedTvShows()
		LoopNumber++
		time.Sleep(48*time.Hour)
	}
}

func GetNumberOfPages() int {
	var err error
	var numberofPages int
	collector := colly.NewCollector()

	collector.OnHTML(".pagination.pagination-lg.justify-content-center", func(element *colly.HTMLElement) {
		element.ForEach(".page-item", func(_ int, element *colly.HTMLElement) {
			title := element.ChildAttr("a", "title")
			href := element.ChildAttr("a", "href")
			if title == "Last" {
				href = strings.ReplaceAll(href, "/tv-show?page=", "")
				numberofPages, err = strconv.Atoi(href)
				HanleError(err)
			}
		})
	})

	collector.Visit("https://tinyzonetv.to/tv-show")

	return numberofPages
}


func JsonMarshal(data interface{}) []byte {
	var buff bytes.Buffer
	encoder := json.NewEncoder(&buff)
	encoder.SetEscapeHTML(false)
	encoder.Encode(data)
	return buff.Bytes()
}


func HanleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}