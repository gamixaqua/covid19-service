package config

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var redisClient *redis.Client

func SetUpRedis() {
	fmt.Println("Starting  Redis")

	redisClient = redis.NewClient(&redis.Options{
		Addr:     "redis-13833.c11.us-east-1-3.ec2.cloud.redislabs.com:13833",
		Password: "UscSNoqLHDp02As8WVTNVChhHWXfTRwI",
		DB:       0,
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		logrus.Fatalf("connection with redis failed with error : %v", err.Error())

	}
	logrus.Info("Redis Connected !")
}

func GetRedisClient() *redis.Client {
	return redisClient
}
