package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"extinctatlas/server/database"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func GenAnswer(w http.ResponseWriter, r *http.Request) {
	// calling function to have database connection
	var coll = database.Connect("creatures");

	godotenv.Load()
	var key string = os.Getenv("KEY")
	// initiate gemini client
	client, err := genai.NewClient(context.TODO(), &genai.ClientConfig{
		APIKey: key,
	})
	if err != nil {
		panic(err)
	}

	info := database.GetInfoCreatures(coll, r.URL.Query().Get("id"))
	strInfo, err := json.Marshal(info)

	question := fmt.Sprintf("You are a helpful and knowledgeable human service agent. Here's some internal background information that might be relevant: %s. A customer has the following question: '%s'. Please answer them directly and conversationally, as a human would. If the background information helps, use it subtly without mentioning it's 'provided data'. If the question goes beyond that, use your general knowledge. Do not reveal you are an AI or a wrapper.", string(strInfo), r.URL.Query().Get("prompt"))

	result, err := client.Models.GenerateContent(context.TODO(), "gemini-2.0-flash", genai.Text(question), nil)

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, result.Text())
}

func GenQuesions(w http.ResponseWriter, r *http.Request) {
	// calling function to have database connection
	var coll = database.Connect("creatures");

	godotenv.Load()
	var key string = os.Getenv("KEY")
	// initiate gemini client
	client, err := genai.NewClient(context.TODO(), &genai.ClientConfig{
		APIKey: key,
	})
	if err != nil {
		panic(err)
	}

	info := database.GetInfoCreatures(coll, r.URL.Query().Get("id"))
	strInfo, err := json.Marshal(info)

	question := fmt.Sprintf("Generate 3 questions based on this information, questions should be short and they are just to demostrate what AI can do, split questions by / sign and don't write any unnecesery infomation like here are your answers or anything like it, also don't write / anything like it at the end of last question, i want straight up answers. here is information, try to make questoins about things that are not straigth up in this information", string(strInfo))
	fmt.Println("questions are generating")
	result, err := client.Models.GenerateContent(context.TODO(), "gemini-2.0-flash", genai.Text(question), nil)

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, result.Text())
}
