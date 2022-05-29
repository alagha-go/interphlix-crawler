package tvshows

import (
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)


func (TvShow *TvShow)SetReleased(element *colly.HTMLElement) {
    layout := "2006-01-02"
    released := element.Text
    released = strings.ReplaceAll(released, "Released: ", "")
    released = strings.ReplaceAll(released, "  ", "")
    released = strings.ReplaceAll(released, "\n", "")
    released = strings.TrimPrefix(released, " ")
    released = strings.TrimSuffix(released, " ")
    Released, _ := time.Parse(layout, released)
    TvShow.Released = &Released
}

func (TvShow *TvShow)SetGenre(element *colly.HTMLElement) {
    element.ForEach("a", func(index int, element *colly.HTMLElement){
        TvShow.Genres = append(TvShow.Genres, element.Text)
    })
}

func (TvShow *TvShow)SetCasts(element *colly.HTMLElement) {
    element.ForEach("a", func(index int, element *colly.HTMLElement){
        TvShow.Casts = append(TvShow.Casts, element.Text)
    })
}

func (TvShow *TvShow)SetDuration(element *colly.HTMLElement) {
    duration := element.Text
    duration = strings.ReplaceAll(duration, "Duration: ", "")
    duration = strings.ReplaceAll(duration, "  ", "")
    duration = strings.ReplaceAll(duration, "\n", "")
    duration = strings.ReplaceAll(duration, "min", "")
    duration = strings.TrimPrefix(duration, " ")
    duration = strings.TrimSuffix(duration, " ")
    if strings.Contains(duration, "N/A") {
        TvShow.Duration = 0
    }else {
        minutes, _ := strconv.Atoi(duration)
        TvShow.Duration = time.Duration(minutes*int(time.Minute))
    }
}

func (TvShow *TvShow)SetCountries(element *colly.HTMLElement) {
    element.ForEach("a", func(index int, element *colly.HTMLElement){
        TvShow.Countries = append(TvShow.Countries, element.Text)
    })
}

func (TvShow *TvShow)SetProducers(element *colly.HTMLElement) {
    production := element.Text
    production = strings.ReplaceAll(production, "Production: ", "")
    production = strings.ReplaceAll(production, "  ", "")
    production = strings.ReplaceAll(production, "\n", "")
    production = strings.TrimPrefix(production, " ")
    production = strings.TrimSuffix(production, " ")
    if production == "N/A" {
        return
    }
    TvShow.Producers = strings.Split(production, ",")
}