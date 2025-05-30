package main

import (
	"log"
	"net/http"

	"extinctatlas/server/handler"
)

func main() {
	const port = "8080"

	http.HandleFunc("/extinctatlas/map", handler.HttpHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
