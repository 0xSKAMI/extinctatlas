package database

import (
	"os"
	"errors"
	"strings"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var coll *mongo.Collection

// function to connect to database
func Connect() (mongo.Client, error) {
	// get env variable and using it connect to database
	godotenv.Load()
	var password string = os.Getenv("PASSWORD")
	var link string = os.Getenv("MONGO_URI");
	client, err := mongo.Connect(options.Client().ApplyURI(strings.Replace(link, "<db_password>", password, 1)));
	// error handling
	if err != nil {
		return *client, errors.New("problem with connect");
	}

	// return coll pointer
	return *client, nil;
}
