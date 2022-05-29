package tvshows

import (
	"strings"

	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (Episode *Episode)SetServers() {
	collector := colly.NewCollector()

    url := "https://tinyzonetv.to/ajax/v2/episode/servers/"+ Episode.Code

    collector.OnHTML(".nav", func(element *colly.HTMLElement) {
        element.ForEach(".nav-item", func(index int, element *colly.HTMLElement) {
            var server Server
			server.ID = primitive.NewObjectID()
            server.WatchID = element.ChildAttr("a", "data-id")
            server.Name = element.ChildAttr("a", "title")
            server.Name = strings.ReplaceAll(server.Name, "Server ", "")
           	Episode.Servers = append(Episode.Servers, server)
        })
    })
    collector.Visit(url)
	Episode.SetID()
	Episode.AddServer()
	Episode.SetServer()
	Episode.SetServer()
}


func (Episode *Episode)SetID() {
	for index, server := range Episode.Servers {
		url := "https://tinyzonetv.to/ajax/get_link/"+ server.WatchID
		data, _, err := GetRequest(url, false)
		HanleError(err)
		res, err := UnmarshalLinkResponse(data)
		HanleError(err)
        if server.Name == "Streamlare" {
			Episode.Servers[index].Id = strings.ReplaceAll(res.Link, "https://streamlare.com/e/", "")
			Episode.Servers[index].Url = "https://streamlare.com/v/" + Episode.Servers[index].Id
		}else if server.Name == "Vidcloud"{
			Episode.Servers[index].Id = strings.ReplaceAll(res.Link, "https://rabbitstream.net/embed-4/", "")
			Episode.Servers[index].Id = strings.ReplaceAll(Episode.Servers[index].Id, "?z=", "")
			Episode.Servers[index].Url = "https://rabbitstream.net/embed/m-download/" + Episode.Servers[index].Id
		}else if server.Name == "UpCloud" {
			Episode.Servers[index].Id = strings.ReplaceAll(res.Link, "https://mzzcloud.life/embed-4/", "")
			Episode.Servers[index].Id = strings.ReplaceAll(Episode.Servers[index].Id, "?z=", "")
			Episode.Servers[index].Url = "https://mzzcloud.life/embed/m-download/" + Episode.Servers[index].Id
		}else {
			Episode.Servers[index].Url = res.Link
		}
	}
}


func (Episode *Episode)AddServer() {
    for _, server := range Episode.Servers {
        if server.Name == "Vidcloud" || server.Name == "UpCloud" {
            collector := colly.NewCollector()

			collector.OnHTML(".download-list", Episode.AddServers)
			collector.Visit(server.Url)
        }
    }
}


func (Episode *Episode)AddServers(element *colly.HTMLElement) {
    element.ForEach(".dl-site", func(_ int, element *colly.HTMLElement) {
		var exist bool = false
		var server Server
		server.ID = primitive.NewObjectID()
		server.Name = element.ChildText(".site-name")
		server.Url = element.ChildAttr("a", "href")
		for index, serve := range Episode.Servers {
			if serve.Name == server.Name {
				Episode.Servers[index].Url = server.Url
				exist = true
			}
		}
		if !exist {
			Episode.Servers = append(Episode.Servers, server)
		}
	})
}



func (Episode *Episode) SetServer() {
	for index := range Episode.Servers {
		if Episode.Servers[index].Name == "Streamlare" {
			Episode.Server = &Episode.Servers[index]
		}
	}
}
