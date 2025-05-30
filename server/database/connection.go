package database

import (
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var coll *mongo.Collection

// function to connect to database
func Connect() mongo.Collection {
	// get env variable and using it connect to database
	godotenv.Load()
	var password string = os.Getenv("PASSWORD")
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb+srv://skami:" + password + "@cluster0.my5ym89.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))
	// error handling
	if err != nil {
		panic(err)
	}

	// creating coll (collection) variable to manage creatures collection
	coll = client.Database("extinctatlas").Collection("creatures")

	// return coll pointer
	return *coll
}