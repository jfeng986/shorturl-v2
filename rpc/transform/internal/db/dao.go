package db

import (
	"context"
	"time"
)

func Put(key, value []byte) error {
	err := redisClient.Set(context.Background(), string(key), string(value), 7*24*time.Hour).Err() // 7 days
	if err != nil {
		return err
	}
	return nil
}

func Get(key []byte) (string, bool) {
	url, err := redisClient.Get(context.Background(), string(key)).Result()
	if err != nil {
		return "", false
	}
	return url, true
}
