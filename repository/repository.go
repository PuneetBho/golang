package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client           *mongo.Client
	db               *mongo.Database
	playerCollection *mongo.Collection
}

func NewRepository(ctx context.Context, dbName string, uri string) (*Repository, error) {
	var err error
	repo := &Repository{}
	opts := options.Client()
	repo.client, err = mongo.Connect(ctx, opts.ApplyURI(uri))

	if err != nil {
		return nil, err
	}

	repo.db = repo.client.Database(dbName)
	repo.playerCollection = repo.db.Collection("gen_players")

	return repo, nil
}

func (r *Repository) Disconnect(ctx context.Context) error {
	return r.client.Disconnect(ctx)
}
