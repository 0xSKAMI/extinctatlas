package database

import (
	"context"

	"extinctatlas/server/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func GetAdrr(coll mongo.Collection) models.Creature {
	data, err := coll.Find(context.TODO(), bson.D{})

	if err != nil {
		panic(err)
	}

	var result models.Creature

	data.All(context.TODO(), &result)

	return result
}

func getInfo() {

}

func getIMG() {

}
