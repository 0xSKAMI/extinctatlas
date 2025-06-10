package main

import (
	"log"
	"net/http"
	"os"

	"extinctatlas/server/handler"

	"github.com/rs/cors"
	"github.com/joho/godotenv"
)

func main() {
	// declaring port
	godotenv.Load()
	var port string = os.Getenv("PORT")

	c := cors.New(cors.Options{
    AllowedOrigins:   []string{"http://localhost:5173"},
    AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete},
    AllowCredentials: true,
	})

	// creating handlers
	http.HandleFunc("/extinctatlas/map", handler.MapHandler)
	http.HandleFunc("/extinctatlas/info/", handler.InfoHandler)
	http.HandleFunc("/extinctatlas/ai/", handler.TestGem)

	// start listening to port
	log.Fatal(http.ListenAndServe(":"+port, c.Handler(http.DefaultServeMux)))
}
