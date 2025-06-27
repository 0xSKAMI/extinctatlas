package main

import (
	"log"
	"net/http"
	"os"

	"extinctatlas/server/handler"

	"github.com/rs/cors"
	"github.com/joho/godotenv"
)

var Collection string 

func main() {
	// declaring port
	godotenv.Load()
	var port string = os.Getenv("PORT")

	c := cors.New(cors.Options{
    AllowedOrigins:   []string{"https://extinctatlas.netlify.app", "https://somethingtesting.duckdns.org", "http://localhost:5173"},
    AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete},
    AllowCredentials: true,
	})

	// creating handlers
	http.HandleFunc("/extinctatlas/map", handler.MapHandler)
	http.HandleFunc("/extinctatlas/info/", handler.InfoHandler)
	http.HandleFunc("/extinctatlas/ai/", handler.GenAnswer)
	http.HandleFunc("/extinctatlas/ai/generate", handler.GenQuesions)

	// start listening to port
	log.Fatal(http.ListenAndServe(":"+port, c.Handler(http.DefaultServeMux)))
}
