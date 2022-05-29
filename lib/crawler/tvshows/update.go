package tvshows


func (TvShow *TvShow) Update() {
	NewTvShow := TvShow
	NewTvShow.Seasons = []Season{}
	NewTvShow.GetSeasons()
	if len(NewTvShow.Seasons) > len(TvShow.Seasons) {
		NewSeasons := NewTvShow.FindNewSeasons(TvShow)
		for index := range NewSeasons {
			TvShow.UpdateSeason(NewSeasons[index])
			TvShow.Seasons = append(TvShow.Seasons, NewSeasons[index])
		}
	}
	NewSeasons := NewTvShow.FindNewEpisodes(TvShow)
	if len(NewSeasons) != 0 {
		for index := range NewSeasons {
			for Eindex := range NewSeasons[index].Episodes {
				NewSeasons[index].Episodes[Eindex].UpdateEpisode(TvShow.ID, NewSeasons[index].ID)
				Sindex := TvShow.FindSeasonIndex(NewSeasons[index].Code)
				if Sindex != -1 {
					TvShow.Seasons[Sindex].Episodes = append(TvShow.Seasons[Sindex].Episodes, NewSeasons[index].Episodes[Eindex])
				}
			}
		}
	}
}

func (TvShow *TvShow) FindNewSeasons(Tvshow *TvShow) []Season {
	var Seasons []Season
	for index := range TvShow.Seasons {
		exists := false
		for index1 := range Tvshow.Seasons {
			if TvShow.Seasons[index].Code == Tvshow.Seasons[index1].Code {
				exists = true
				break
			}
		}
		if !exists {
			Seasons = append(Seasons, TvShow.Seasons[index])
		}
	}
	return Seasons
}


func (TvShow *TvShow) FindNewEpisodes(Tvshow *TvShow) []Season {
	var Seasons []Season
	for index := range TvShow.Seasons {
		var NewSeason Season
		Season := Tvshow.FindSeason(TvShow.Seasons[index].Code)
		NewSeason = Season
		NewSeason.Episodes = []Episode{}
		for Eindex := range TvShow.Seasons[index].Episodes {
			exists := false
			for Eindex1 := range Season.Episodes {
				if TvShow.Seasons[index].Episodes[Eindex].Code == Season.Episodes[Eindex1].Code {
					exists = true
					break
				}
			}
			if !exists {
				NewSeason.Episodes = append(NewSeason.Episodes, TvShow.Seasons[index].Episodes[Eindex])
			}
		}
		if len(NewSeason.Episodes) > 0 {
			Seasons = append(Seasons, NewSeason)
		}
	}
	return Seasons
}