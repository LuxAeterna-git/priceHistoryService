package repository

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	db            *mongo.Collection
	lastPricesReq primitive.ObjectID // id of last requested raw data for aware of duplicate data
}

func NewMongo(ctx context.Context, host, port, username, password, database, collection, authDB string) (*MongoDB, error) {
	mongoURL := fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)

	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		AuthSource: authDB,
		Username:   username,
		Password:   password,
	})

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect with mongo")
	}
	if err = client.Ping(ctx, nil); err != nil {
		return nil, errors.Wrap(err, "Failed to ping mongo")
	}
	return &MongoDB{
		db: client.Database(database).Collection(collection),
	}, nil

}
