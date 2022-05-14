package tvshows

import (
	"fmt"

	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func (TvShow *Movie)GetSeasons() {
	collector := colly.NewCollector()
	url := "https://tinyzonetv.to/ajax/v2/tv/seasons/" + TvShow.Code

	collector.OnHTML(".dropdown-menu.dropdown-menu-new", TvShow.CollectAllSeasons)

	collector.Visit(url)
}

func (TvShow *Movie)CollectAllSeasons(element *colly.HTMLElement) {
	element.ForEach("a", func(index int, element *colly.HTMLElement) {
		var Season Season
		Season.ID = primitive.NewObjectID()
		Season.Index = index
		Season.Code = element.Attr("data-id")
		Season.Name = element.Text
		Season.GetEpisodes()
		TvShow.Seasons = append(TvShow.Seasons, Season)
	})
}


func (TvShow *Movie) SeasonExist(Season Season) bool {
	for index := range TvShow.Seasons {
		if Season.Code == TvShow.Seasons[index].Code {
			return true
		}
	}
	return false
}

func (TvShow *Movie) UpdateSeason(Season Season) {
	url := fmt.Sprintf("https://s1.interphlix.com/movies/%s/updateseason", TvShow.ID.Hex())
	body := JsonMarshal(Season)
	PostRequest(url, body, false)
}