package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDbInit(ctx context.Context, host, port, dbName string) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(
		fmt.Sprintf(
			"mongodb://%s:%s/",
			host, port,
		),
	)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client.Database(dbName), nil
}
