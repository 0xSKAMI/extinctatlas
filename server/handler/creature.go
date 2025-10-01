package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"extinctatlas/server/database"
)


// function to return map parameters
func MapHandler(w http.ResponseWriter, r *http.Request) {
	// calling function to have database connection
	var client, err = database.Connect();
	//error handling for connection
	if err != nil {
		log.Printf("MapHandler: db connect error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// creating coll (collection) variable to manage creatures collection
	var coll = client.Database("extinctatlas").Collection("creatures")

	// get all adresses and transform them to JSON bytes
	result, err := database.GetAdrrCreatures(*coll)
	//error handling for address query
	if err != nil {
		log.Printf("MapHandler: db query error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	result2, err := json.Marshal(result)
	// error handling
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// return JSON andswer
	fmt.Fprint(w, string(result2))
}

// function to return info about one creature
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	// calling function to have database connection
	var client, err = database.Connect();
	//error handling for connection
	if err != nil {
		log.Printf("InfoHandler: db connect error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// creating coll (collection) variable to manage creatures collection
	var coll = client.Database("extinctatlas").Collection("creatures")

	// get url and extract id from it
	basePath := r.URL.Path
	pathArray := strings.Split(basePath, "/")	
	// erro handling
	if len(pathArray) > 3 {
		http.NotFoundHandler()
	}

	// get info about one creature (using id) and transforming it to JSON bytes
	result, err := database.GetInfoCreatures(*coll, pathArray[3])
	//error handling for query 
	if err != nil {
		log.Printf("InfoHandler: db query error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	result2, err := json.Marshal(result)
	// error handling
	if err != nil {
		log.Printf("InfoHandler: response encode error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// return JSON format answer
	fmt.Fprint(w, string(result2))
}
