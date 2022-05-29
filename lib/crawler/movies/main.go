package movies

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)


var (
	Movies []Movie
	LoopNumber int
	TotalNumberOfPages int
	CurrentPageNumber int
	CurrentPageCollectedMovies int
	TotalNumberOfMovies int
	CurrentMovie	int
)


func Main() {
	for {
		LoopNumber++
		LoadMovies()
		CollectPages(GetNumberOfPages())
		UploadUnUploadedMovies()
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
				href = strings.ReplaceAll(href, "/movie?page=", "")
				numberofPages, err = strconv.Atoi(href)
				HanleError(err)
			}
		})
	})

	collector.Visit("https://tinyzonetv.to/movie")

	return numberofPages
}


func LoadMovies() {
	data, err := ioutil.ReadFile("./DB/Movies/movies.json")
	HanleError(err)
	json.Unmarshal(data, &Movies)
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