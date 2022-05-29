package tvshows

import (
	"fmt"

	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func (TvShow *TvShow)GetSeasons() {
	collector := colly.NewCollector()
	url := "https://tinyzonetv.to/ajax/v2/tv/seasons/" + TvShow.Code

	collector.OnHTML(".dropdown-menu.dropdown-menu-new", TvShow.CollectAllSeasons)

	collector.Visit(url)
}

func (TvShow *TvShow)CollectAllSeasons(element *colly.HTMLElement) {
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


func (TvShow *TvShow) FindSeason(code string) Season {
	for index := range TvShow.Seasons {
		if TvShow.Seasons[index].Code == code {
			return TvShow.Seasons[index]
		}
	}
	return Season{}
}

func (TvShow *TvShow) FindSeasonIndex(code string) int {
	for index := range TvShow.Seasons {
		if TvShow.Seasons[index].Code == code {
			return index
		}
	}
	return -1
}


func (TvShow *TvShow) UpdateSeason(Season Season) {
	url := fmt.Sprintf("https://s1.interphlix.com/tv-shows/%s/addseason", TvShow.ID.Hex())
	body := JsonMarshal(Season)
	PostRequest(url, body, false)
}