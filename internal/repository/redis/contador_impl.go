package redis

import (
	"context"
	"encurtador-url-go/internal/database"
)

type ContadorImpl struct{}

func NewContador() *ContadorImpl {
	return &ContadorImpl{}
}

func (contador *ContadorImpl) Next() (int64, error) {
	return database.RedisConnection.Incr(context.Background(), database.KeyContador).Result()
}
