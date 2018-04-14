package main

import (
	"net/http"
	"log"
	"encoding/json"
	"github.com/krushelnytskyi/contact-vitalii/api"
)

func main () {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			json.NewEncoder(writer).Encode(api.Authorize(request.RemoteAddr))
		}
	})

	log.Fatal(http.ListenAndServe(":9000", nil))
}


