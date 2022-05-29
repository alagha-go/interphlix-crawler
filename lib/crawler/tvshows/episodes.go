package tvshows

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func (Season *Season)GetEpisodes() {
	collector := colly.NewCollector()
	url := "https://tinyzonetv.to/ajax/v2/season/episodes/" + Season.Code

	collector.OnHTML(".nav", Season.CollectAllEpisodes)

	collector.Visit(url)
}

func (Season *Season)CollectAllEpisodes(element *colly.HTMLElement) {
	element.ForEach(".nav-item", func(_ int, element *colly.HTMLElement) {
		var Episode Episode
		Episode.ID = primitive.NewObjectID()
		Episode.Name = element.ChildText("a")
		index := strings.Index(Episode.Name, "Eps")
		end := strings.Index(Episode.Name, "\n")
		Episode.Index, _ = strconv.Atoi(Episode.Name[index+4:end])
		index = strings.Index(Episode.Name, "\n")
		Episode.Name = Episode.Name[index+1:]
		Episode.Name = strings.ReplaceAll(Episode.Name, "                    : ", "")
		Episode.Code = element.ChildAttr("a", "data-id")
		Episode.SetServers()
		Season.Episodes = append(Season.Episodes, Episode)
		EpisodesLength = EpisodesLength+len(Season.Episodes)
	})
}

func (Episode *Episode) UpdateEpisode(MovieID, SeasonID primitive.ObjectID) {
	url := fmt.Sprintf("https://s1.interphlix.com/tv-shows/%s/%s/addepisode", MovieID.Hex(), SeasonID)
	body := JsonMarshal(Episode)
	PostRequest(url, body, false)
}