package handler

import "net/http"


func Main() {
	http.HandleFunc("/", GetStats)
}

func GetStats(res http.ResponseWriter, req *http.Request) {
	
}