package datasource

import (
	"context"
	"template-ulamm-backend-go/utils"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB() (*mongo.Database, error) {
	mongoOptions := options.Client().ApplyURI(utils.GetConfig().Mongo.Host)
	cmdMonitor := &event.CommandMonitor{
		Started: func(ctx context.Context, cse *event.CommandStartedEvent) {
			utils.GetLogger().Debug(cse.Command.String())
		},
	}

	mongoOptions.SetMonitor(cmdMonitor)

	mongoClient, err := mongo.Connect(context.TODO(), mongoOptions)
	if err != nil {
		return nil, err
	}

	return mongoClient.Database(utils.GetConfig().Mongo.DbName), nil
}
