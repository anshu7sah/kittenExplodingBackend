package database

import(
	"context"
	"github.com/go-redis/redis/v8"
	"os"
)

var Ctx=context.Background()

func CreateClient(dbNo int) *redis.Client{
	rdb:=redis.NewClient(&redis.Options{
		Addr:os.Getenv("DB_ADDR"),
		Username:os.Getenv("DB_USER"),
		Password:os.Getenv("DB_PASS"),
		DB:dbNo,
	})
	return rdb
}