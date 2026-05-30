package database

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var RedisConnection *redis.Client
var _ValorInicialDoContador = 999999
var KeyContador = "contador_urls"

func ConnectionRedis() {
	RedisConnection = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
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
