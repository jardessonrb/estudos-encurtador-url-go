package database

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisConnection *redis.Client
var _ValorInicialDoContador = 999999
var KeyContador = "contador_urls"

func ConnectionRedis() {
	REDIS_HOST := os.Getenv("REDIS_HOST")
	REDIS_PORT := os.Getenv("REDIS_PORT")

	RedisConnection = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", REDIS_HOST, REDIS_PORT),
	})

	exists, err := RedisConnection.Exists(context.Background(), KeyContador).Result()

	if err != nil {
		panic(err)
	}

	if exists == 0 {
		err := RedisConnection.Set(
			context.Background(),
			"contador_urls",
			_ValorInicialDoContador,
			0,
		).Err()

		if err != nil {
			panic(err)
		}
	}

}
