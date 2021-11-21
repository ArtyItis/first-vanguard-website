package main

import (
	"forgottennw/app/route"
	"log"
	"net/http"
)

func main() {
	route.MapURLs()
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
