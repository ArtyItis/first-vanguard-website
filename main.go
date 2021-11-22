package main

import (
	"forgottennw/app/route"
	"log"
	"net/http"
)

func main() {
	router := route.NewRouter()
	log.Println("Listening...")
	http.ListenAndServe(":8080", router)
}
