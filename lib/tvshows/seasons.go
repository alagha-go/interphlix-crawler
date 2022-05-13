package tvshows

import (

	"github.com/gocolly/colly"
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
		Season.Index = index
		Season.Code = element.Attr("data-id")
		Season.Name = element.Text
		Season.GetEpisodes()
		TvShow.Seasons = append(TvShow.Seasons, Season)
	})
	// TvShow.Seasons = seasons
	// PrintCyan(len(seasons))
	// PrintCyan(len(TvShow.Seasons))
}
