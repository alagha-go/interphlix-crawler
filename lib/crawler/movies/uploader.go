package movies

import (
	"encoding/json"
	"io/ioutil"
)


func UploadUnUploadedMovies() {
	var Movies []Movie
	data, err := ioutil.ReadFile("./DB/Movies/movies.json")
	HanleError(err)
	json.Unmarshal(data, &Movies)
	for index := range Movies {
		if !Movies[index].Uploaded && Movies[index].Collected{
			Movies[index].Upload()
			if !Movies[index].Uploaded {
				Movies[index].SetMovieID()
				Movies[index].Upload()
			}
			ioutil.WriteFile("./DB/Movies/movies.json", JsonMarshal(Movies), 0755)
		}
	}
}