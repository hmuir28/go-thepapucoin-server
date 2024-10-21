package database

import (
    "context"
    "fmt"
    "github.com/redis/go-redis/v9"
    "log"
)

func NewRedisClient() *redis.Client {
	// password := os.Getenv("REDIS_PASSWORD")
	// database := os.Getenv("REDIS_DB")
	// address := os.Getenv("REDIS_ADDRESS")

	password := "admin"
	database := 0
	address  := "localhost:6379"

    client := redis.NewClient(&redis.Options{
        Addr:     address,
        Password: password,
        DB:       database,
    })

    // Check if the connection is successful
    _, err := client.Ping(context.Background()).Result()

	if err != nil {
        log.Fatalf("Could not connect to Redis: %v", err)
    }

    fmt.Println("Connected to Redis!")
    return client
}
