package tvshows

import (
	"encoding/json"
	"strings"

	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



func (TvShow *TvShow)CollectTvShowContent() {
    collector := colly.NewCollector()

    collector.OnHTML(".description", func(element *colly.HTMLElement){
        TvShow.Description = element.Text
        TvShow.Description = strings.ReplaceAll(TvShow.Description, "\n", "")
        TvShow.Description = strings.ReplaceAll(TvShow.Description, "  ", "")
        TvShow.Description = strings.TrimPrefix(TvShow.Description, " ")
        TvShow.Description = strings.TrimSuffix(TvShow.Description, " ")
    })

    
    collector.OnHTML(".elements", TvShow.SetElements)
    
    collector.Visit(TvShow.PageUrl)
}


func (TvShow *TvShow)SetElements(element *colly.HTMLElement) {
    functions := []func(element *colly.HTMLElement){}
    functions = append(functions,  TvShow.SetReleased)
    functions = append(functions,  TvShow.SetGenre)
    functions = append(functions,  TvShow.SetCasts)
    functions = append(functions,  TvShow.SetDuration)
    functions = append(functions,  TvShow.SetCountries)
    functions = append(functions,  TvShow.SetProducers)
    element.ForEach(".row-line", func(index int, element *colly.HTMLElement){
        functions[index](element)
    })
}



func (TvShow *TvShow) SetTvShowID() {
    ID := primitive.NewObjectID()
    for _, movie := range TvShows {
        if ID == movie.ID {
            TvShow.SetTvShowID()
        }
    }
    TvShow.ID = ID
}

func (TvShow *TvShow) Exists() bool {
	for index := range TvShows {
		if TvShows[index].Code == TvShow.Code {
			return true
		}
	}
	return false
}


func (Tvshow *TvShow) Upload() {
	var newTvShow TvShow
	data, _, _ := PostRequest("https://s1.interphlix.com/movies/upload", JsonMarshal(Tvshow), false)
	err := json.Unmarshal(data, &newTvShow)
	if err != nil {
		return
	}
	Tvshow.Uploaded = true
}