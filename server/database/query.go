package database

import (
	"context"
	"fmt"

	"extinctatlas/server/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// function to get all of the addresses in the db
func GetAdrr(coll mongo.Collection) []models.Creature {
	// getting all of the adresses in the db and saving them in data (encoded)
	data, err := coll.Find(context.TODO(), bson.D{}, options.Find().SetProjection(bson.D{{"_id", 1}, {"name", 1}, {"coordinates", 1}}))
	// error handling
	if err != nil {
		panic(err)
	}

	// creating result (value we return) and saving decoded data there
	var result []models.Creature
	err = data.All(context.TODO(), &result)
	// error handling 
	if (err != nil) {
		fmt.Println(err)
		panic(err)
	}
	
	// return result
	return result
}

// function to get information about one creatur
func GetInfo(coll mongo.Collection, ID string) models.Creature {
	// getting encoded information, decode it and save it in data
	var data mongo.SingleResult
	err := coll.FindOne(context.TODO(), bson.D{{"_id", ID}}).Decode(&data)
	// error handling
	if err != nil {
		panic(err)
	}

	// create result variable and save array of decoded info in it
	var result models.Creature
	data.Decode(&result)
	// error handling
	if (err != nil) {
		fmt.Println(err)
		panic(err)
	}

	// return result
	return result
}

// it does not serve any purpose for now
func getIMG() {
}
