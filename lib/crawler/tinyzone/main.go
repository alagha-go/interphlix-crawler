package tinyzone

import (
	"crawler/lib/crawler/tinyzone/movies"
	"crawler/lib/crawler/tinyzone/tvshows"
)


func Main() {
	go movies.Main()
	go tvshows.Main()
}