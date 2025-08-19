package database

import (
	"context"
	"log"
	"errors"

	"extinctatlas/server/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// function to get all of the addresses in the db
func GetAdrrCreatures(coll mongo.Collection) ([]models.Creature, error) {
	// creating result (value we return) and saving decoded data there
	var result []models.Creature

	// getting all of the adresses in the db and saving them in data (encoded)
	data, err := coll.Find(context.TODO(), bson.D{}, options.Find().SetProjection(bson.D{{"_id", 1}, {"name", 1}, {"coordinates", 1}, {"imageurl", 1}}))
	// error handling
	if err != nil {
		log.Println("problem with query: %v", err);
		return result, errors.New("problem with query");
	}

	err = data.All(context.TODO(), &result)
	// error handling 
	if (err != nil) {
		log.Println("problem with tranforming addresses into json: %v", err);
		return result, errors.New("problem with query");
	}
	
	// return result
	return result, nil
}

// function to get information about one creatur
func GetInfoCreatures(coll mongo.Collection, ID string) (models.Creature, error) {
	//transforming ID string to objectid
	_id, err := primitive.ObjectIDFromHex(ID);
	// getting encoded information, decode it and save it in data
	var result models.Creature;
	err = coll.FindOne(context.TODO(), bson.D{{"_id", _id}}).Decode(&result);
	// error handling
	if err != nil {
		if (err == mongo.ErrNoDocuments) {
		}
		log.Println("GetInfoCreatures: %v", err);
		return result, errors.New("getting creature info");
	}

	// create result variable and save array of decoded info in it
	return result, nil;
}
