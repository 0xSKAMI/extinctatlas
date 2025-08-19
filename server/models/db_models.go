package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// defining Creatures structure which is used in mongodb database
type Creature struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string
	Type        string
	LastSeen   	int 
	Reason      string
	HeightCM    int
	WeightKG    float64
	Diet        []string
	ImageURL    string
	Polygon 		bson.M `bson:"polygon"` // GeoJSON Polygon
}
