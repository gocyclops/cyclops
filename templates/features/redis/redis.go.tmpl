package myredis

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

// Initializa a redis connection
// InitRedis initializes a connection to a Redis server using the default address
// "127.0.0.1:6379". It creates a new Redis client with the specified options and
// tests the connection by sending a PING command. If the connection is successful,
// it prints "Connected to Redis!". If the connection fails, it prints an error message.
func InitRedis() {
	redisAddress := os.Getenv("REDIS_URL")

	options := &redis.Options{
		Addr: redisAddress,
	}

	RedisClient = redis.NewClient(options)

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Could not connect to Redis:", err)
	} else {
		fmt.Println("Connected to Redis!")
	}
}
