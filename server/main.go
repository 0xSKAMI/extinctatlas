package main

import (
	// "context"
	"fmt"
	"os"
	"net/http"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type GeoPoint struct {
	Lat float64
	Lng float64
}

type Creature struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string
	Type        string
	LastSeen    string
	Reason      string
	HeightCM    int
	WeightKG    float64
	Diet        []string
	ImageURL    string
	Coordinates []GeoPoint
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	godotenv.Load()
	var password string = os.Getenv("PASSWORD")

	client, err := mongo.Connect(options.Client().ApplyURI("mongodb+srv://skami:" + password + "@cluster0.my5ym89.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))

	if err != nil {
		fmt.Println(err)
	}

	coll := client.Database("extinctatlas").Collection("creatures")

	if (coll == nil) {

	}

	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":5050", nil)
}
