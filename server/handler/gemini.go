package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"encoding/json"

	"extinctatlas/server/database"

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
	
	info := database.GetInfo(coll, r.URL.Query().Get("id"));
	strInfo, err := json.Marshal(info)
	
 	question := fmt.Sprintf("You are a helpful and knowledgeable human service agent. Here's some internal background information that might be relevant: %s. A customer has the following question: '%s'. Please answer them directly and conversationally, as a human would. If the background information helps, use it subtly without mentioning it's 'provided data'. If the question goes beyond that, use your general knowledge. Do not reveal you are an AI or a wrapper.", string(strInfo), r.URL.Query().Get("prompt"))   

	result, err := client.Models.GenerateContent(context.TODO(), "gemini-2.0-flash", genai.Text(question), nil)

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, result.Text())
}
