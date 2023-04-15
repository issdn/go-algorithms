package main

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func GetRedisClient() *redis.Client {
	lerr := godotenv.Load()
	if lerr != nil {
		panic(lerr)
	}

	DB := os.Getenv("REDIS_DB")
	intDB, convErr := strconv.Atoi(DB)

	if convErr != nil {
		panic(convErr)
	}

	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       intDB,
	})
}
