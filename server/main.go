package main

import (
	// "context"
	"fmt"
	"net/http"

	"extinctatlas/server/database"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	// listen to port
	http.ListenAndServe(":5050", nil)

	var coll = database.Connect()

	result := database.GetAdrr(coll)

	fmt.Println(result)
}
