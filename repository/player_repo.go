package repository

import (
	"context"
	"go-mongo/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) GetPlayers(ctx context.Context, position string, salary int64, enable bool, sortField string) (*[]model.Player, error) {

	filters := bson.M{}

	if position != "" {
		filters["position"] = position
	}

	if salary > 0 {
		filters["salary"] = salary
	}

	if enable {
		filters["isEnabled"] = enable
	}

	findOptions := options.Find()

	if sortField != "" {
		findOptions.SetSort(bson.M{sortField: 1})
	}

	log.Println(filters)
	players := []model.Player{}
	cursor, err := r.playerCollection.Find(ctx, filters, findOptions)
	defer cursor.Close(ctx)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var player model.Player
		if err := cursor.Decode(&player); err != nil {
			return nil, err
		}
		players = append(players, player)
	}

	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return &players, nil
}

func (r *Repository) AddPlayer(ctx context.Context, player model.Player) error {
	_, err := r.playerCollection.InsertOne(ctx, player)
	return err
}
