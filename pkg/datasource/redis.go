package datasource

import (
	"context"
	"template-ulamm-backend-go/utils"

	"github.com/redis/go-redis/v9"
)

func NewRedis() (*redis.Client, error) {
	// Redis Database
	rdb := redis.NewClient(&redis.Options{
		Addr:     utils.GetConfig().Redis.Host,
		Password: utils.GetConfig().Redis.Password, // no password set
		DB:       0,                                // use default DB
	})

	status := rdb.Ping(context.Background())
	if err := status.Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
