package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"extinctatlas/server/database"
)

// calling function to have database connection
var coll = database.Connect()

// function to return map parameters
func MapHandler(w http.ResponseWriter, r *http.Request) {
	// get all adresses and transform them to JSON bytes
	var result = database.GetAdrr(coll)
	result2, err := json.Marshal(result)
	// error handling
	if err != nil {
		panic(err)
	}

	// return JSON andswer
	fmt.Fprint(w, string(result2))
}

// function to return info about one creature
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	// get url and extract id from it
	basePath := r.URL.Path
	pathArray := strings.Split(basePath, "/")	
	// erro handling
	if len(pathArray) > 3 {
		http.NotFoundHandler()
	}

	// get info about one creature (using id) and transforming it to JSON bytes
	var result = database.GetInfo(coll, pathArray[3])
	result2, err := json.Marshal(result)
	// error handling
	if err != nil {
		panic(err)
	}

	// return JSON format answer
	fmt.Fprint(w, string(result2))
}
