package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
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