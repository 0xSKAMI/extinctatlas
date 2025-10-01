package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"log"

	"extinctatlas/server/database"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func GenAnswer(w http.ResponseWriter, r *http.Request) {
	// calling function to have database connection
	var client, err = database.Connect()
	//error handling for connection
	if err != nil {
		log.Printf("GetAnswer: db connect error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// creating coll (collection) variable to manage creatures collection
	var coll = client.Database("extinctatlas").Collection("creatures")

	godotenv.Load()
	var key string = os.Getenv("KEY")
	// initiate gemini client
	gemini_client, err := genai.NewClient(context.TODO(), &genai.ClientConfig{
		APIKey: key,
	})
	//error handling for gemini client init
	if err != nil {
		log.Printf("GetAnser: gemini client init error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	info, err := database.GetInfoCreatures(*coll, r.URL.Query().Get("id"))
	//error handling for query 
	if err != nil {
		log.Printf("GetAnswer: db query error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	strInfo, err := json.Marshal(info)
	// error handling
	if err != nil {
		log.Printf("GetAnswer: response encode error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	question := fmt.Sprintf("You are a helpful and knowledgeable human service agent. Here's some internal background information that might be relevant: %s. A customer has the following question: '%s'. Please answer them directly and conversationally, as a human would. If the background information helps, use it subtly without mentioning it's 'provided data'. If the question goes beyond that, use your general knowledge. Do not reveal you are an AI or a wrapper.", string(strInfo), r.URL.Query().Get("prompt"))

	result, err := gemini_client.Models.GenerateContent(context.TODO(), "gemini-2.0-flash", genai.Text(question), nil)

	if err != nil {
		log.Printf("GetAnswer: context generating error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, result.Text())
}

func GenQuesions(w http.ResponseWriter, r *http.Request) {
	// calling function to have database connection
	var client, err = database.Connect();
	//error handling for connection
	if err != nil {
		log.Printf("GenQuestions: db connect error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// creating coll (collection) variable to manage creatures collection
	var coll = client.Database("extinctatlas").Collection("creatures")

	godotenv.Load()
	var key string = os.Getenv("KEY")
	// initiate gemini client
	gemini_client, err := genai.NewClient(context.TODO(), &genai.ClientConfig{
		APIKey: key,
	})
	//error handling for gemini client init
	if err != nil {
		log.Printf("GenQuestions: gemini client init error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	info, err := database.GetInfoCreatures(*coll, r.URL.Query().Get("id"))
	//error handling for query 
	if err != nil {
		log.Printf("GenQeustions: db query error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	strInfo, err := json.Marshal(info)
	// error handling
	if err != nil {
		log.Printf("GenQuestions: response encode error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	question := fmt.Sprintf("Generate 3 questions based on this information, questions should be short and they are just to demostrate what AI can do, split questions by / sign and don't write any unnecesery infomation like here are your answers or anything like it, also don't write / anything like it at the end of last question, i want straight up answers. here is information, try to make questoins about things that are not straigth up in this information", string(strInfo))
	fmt.Println("questions are generating")
	result, err := gemini_client.Models.GenerateContent(context.TODO(), "gemini-2.0-flash", genai.Text(question), nil)

	if err != nil {
		log.Printf("GenQuestions: context generating error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, result.Text())
}
