package main

import (
	"crawler/lib/handler"
	"net/http"
)

var (
	PORT = ":7000"
)

func main() {
	handler.Main()
	http.ListenAndServe(PORT, nil)
}