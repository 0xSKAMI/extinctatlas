package main

import (
	"log"
	"net/http"

	"extinctatlas/server/handler"
)

func main() {
 // declaring port 
	const port = "8080"

	// creating handlers
	http.HandleFunc("/extinctatlas/map", handler.MapHandler)
	http.HandleFunc("/extinctatlas/info/", handler.InfoHandler)
	http.HandleFunc("/extinctatlas/ai", handler.TestGem);

	// start listening to port
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
