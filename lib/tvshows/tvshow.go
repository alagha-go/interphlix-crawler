package tvshows

import (
	"strings"

	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson/primitive"
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


func (TvShow *Movie) SetMovieID() {
    ID := primitive.NewObjectID()
    for _, movie := range TvShows {
        if ID == movie.ID {
            TvShow.SetMovieID()
        }
    }
    TvShow.ID = ID
}


func (TvShow *Movie) CheckUpdate() {
    Tvshow := FindTvShow(TvShow.ID)
    if TvShow.ID != Tvshow.ID {
        return
    }
    for Sindex := range TvShow.Seasons {
        if !Tvshow.SeasonExist(TvShow.Seasons[Sindex]) {
            Tvshow.UpdateSeason(TvShow.Seasons[Sindex])
        }else {
            for index := range TvShow.Seasons[Sindex].Episodes {
                if !Tvshow.Seasons[index].EpisodeExist(TvShow.Seasons[Sindex].Episodes[index]) {
                    Tvshow.Seasons[Sindex].Episodes[index].UpdateEpisode(Tvshow.ID, Tvshow.Seasons[index].ID)
                }
            }
        }
    }
}

func FindTvShow(ID primitive.ObjectID) Movie {
    for _, TvShow := range DBTvShows {
        if TvShow.ID.Hex() == ID.Hex() {
            return TvShow
        }
    }
    return Movie{}
}