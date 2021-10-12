package Lib

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func RedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return client
}

func RedisSet(rc *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("redis", rc)
		c.Next()
	}
}

func Set(client *redis.Client, key, value string) error {
	err := client.Set(key, value, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func Get(client *redis.Client, key string) (string, error) {
	value, err := client.Get(key).Result()
	if err == redis.Nil {
		return "no value found", nil
	} else if err != nil {
		return "", err
	}

	return value, nil
}
