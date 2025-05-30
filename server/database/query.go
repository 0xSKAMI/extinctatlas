package database

import (
	"context"
	"fmt"

	"extinctatlas/server/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func GetAdrr(coll mongo.Collection) []models.Creature {
	data, err := coll.Find(context.TODO(), bson.D{}, options.Find().SetProjection(bson.D{{"_id", 1}, {"name", 1}, {"coordinates", 1}}))

	if err != nil {
		panic(err)
	}

	var result []models.Creature

	err = data.All(context.TODO(), &result)

	if (err != nil) {
		fmt.Println(err)
		panic(err)
	}

	return result
}

func getInfo(ID string) models.Creature {
	var data mongo.SingleResult
	
	err := coll.FindOne(context.TODO(), bson.D{{"_id", ID}}).Decode(&data)

	if err != nil {
		panic(err)
	}

	var result models.Creature
	
	data.Decode(&result)
	
	if (err != nil) {
		fmt.Println(err)
		panic(err)
	}

	return result
}

func getIMG() []models.Creature {
	data, err := coll.Find(context.TODO(), bson.D{}, options.Find().SetProjection(bson.D{{"_id", 1}, {"name", 1}, {"coordinates", 1}}))

	if err != nil {
		panic(err)
	}

	var result []models.Creature

	err = data.All(context.TODO(), &result)

	if (err != nil) {
		fmt.Println(err)
		panic(err)
	}

	return result
}
