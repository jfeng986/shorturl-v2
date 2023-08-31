package db

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"gopkg.in/ini.v1"
)

var (
	rdbHost     string
	rdbPort     string
	rdbPass     string
	redisClient *redis.Client
)

func init() {
	file, err := ini.Load("./internal/db/config.ini")
	if err != nil {
		log.Println("Configuration file reading error. Please check the file path: ", err)
		log.Println(err)
	}
	LoadRedis(file)
	ctx := context.Background()

	redisClient = redis.NewClient(&redis.Options{
		Addr:     rdbHost + ":" + rdbPort,
		Password: rdbPass,
		DB:       0,
	})

	_, err = redisClient.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	log.Println("Redis connected successfully")
}

func LoadRedis(file *ini.File) {
	rdbHost = file.Section("redis").Key("DB_HOST").String()
	rdbPort = file.Section("redis").Key("DB_PORT").String()
	rdbPass = file.Section("redis").Key("DB_PASS").String()
}
