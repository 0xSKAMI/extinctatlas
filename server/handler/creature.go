package handler

import (
	"fmt"
	"net/http"
	"encoding/json"

	"extinctatlas/server/database"
)

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	var coll = database.Connect()

	result := database.GetAdrr(coll)

	fmt.Println(result)

	test, err := json.Marshal(result)

	if (err != nil) {
		panic(err)
	}

	fmt.Fprint(w, string(test))
}