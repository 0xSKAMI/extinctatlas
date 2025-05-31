package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func TestGem(w http.ResponseWriter, r *http.Request) {
	godotenv.Load()
	var key string = os.Getenv("KEY")
	// initiate gemini client
	client, err := genai.NewClient(context.TODO(), &genai.ClientConfig{
		APIKey: key,
	})
	if err != nil {
		panic(err)
	}

	result, err := client.Models.GenerateContent(context.TODO(), "gemini-2.0-flash", genai.Text(r.URL.Query().Get("prompt")), nil)

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, result.Text())
}
