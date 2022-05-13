package tvshows

import (
	"strings"

	"github.com/gocolly/colly"
)


func (TvShow *Movie)CollectTvShowContent() {
    collector := colly.NewCollector()

    collector.OnHTML(".description", func(element *colly.HTMLElement){
        TvShow.Description = element.Text
        TvShow.Description = strings.ReplaceAll(TvShow.Description, "\n", "")
        TvShow.Description = strings.ReplaceAll(TvShow.Description, "  ", "")
        TvShow.Description = strings.TrimPrefix(TvShow.Description, " ")
        TvShow.Description = strings.TrimSuffix(TvShow.Description, " ")
    })

    
    collector.OnHTML(".elements", TvShow.SetElements)
	TvShow.GetSeasons()
    
    collector.Visit(TvShow.PageUrl)
}


func (TvShow *Movie)SetElements(element *colly.HTMLElement) {
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