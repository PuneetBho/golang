package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Player struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	PlayerName string             `bson:"playerName"`
	Position   string             `bson:"position"`
	IsEnabled  bool               `bson:"isEnabled"`
	Salary     int64              `bson:"salary"`
}
