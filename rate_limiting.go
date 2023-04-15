package main

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type RateLimitError struct {
	message string
	code    int
}

func (e *RateLimitError) Error() string {
	return e.message
}

type SlidingWindow struct {
	Interval    int64
	Treshold    int64
	redisClient redis.Client
	ctx         context.Context
}

func NewSlidingWindow(interval int64, treshold int64) *SlidingWindow {
	envErr := godotenv.Load()
	if envErr != nil {
		panic(envErr)
	}
	redisDb, envDbErr := strconv.Atoi(os.Getenv("REDIS_DB"))
	if envDbErr != nil {
		panic(envDbErr)
	}
	limiter := &SlidingWindow{
		Interval: interval,
		Treshold: treshold,
		redisClient: *redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_HOST"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       redisDb,
		}),
		ctx: context.Background(),
	}
	limiter.redisClient.FlushAll(limiter.ctx)
	return limiter
}

func (fw *SlidingWindow) GetNumberOfRequests(IP string) (int64, error) {
	numberOfRequests, err := fw.redisClient.LLen(fw.ctx, IP).Result()
	if err == redis.Nil {
		return 0, redis.Nil
	} else if err != nil {
		return 0, err
	} else {
		return numberOfRequests, nil
	}
}

func (fw *SlidingWindow) AddCurrentTimestamp(IP string) error {
	now := time.Now().Unix()
	fw.redisClient.LRem(fw.ctx, IP, -1, now-int64(fw.Interval*60))
	return fw.redisClient.RPush(fw.ctx, IP, now).Err()
}

func (fw *SlidingWindow) Request(IP string) (int64, error) {
	numberOfRequests, err := fw.GetNumberOfRequests(IP)
	if err == redis.Nil {
		err := fw.AddCurrentTimestamp(IP)
		if err != nil {
			return -1, err
		}
	} else if err != nil {
		return -1, err
	}

	if numberOfRequests >= fw.Treshold {
		return -1, &RateLimitError{"too many requests", 429}
	} else {
		err := fw.AddCurrentTimestamp(IP)
		if err != nil {
			return -1, err
		} else {
			return int64(numberOfRequests) + 1, nil
		}
	}
}
