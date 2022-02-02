package main

import (
	"go-match/cmd/http"
)

type IdentifyRequest struct {
	RequestData map[string]interface{} `json:"requestData"`
}

func main() {
	app := http.NewApp()
	app.Start()
}
