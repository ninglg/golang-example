package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	fmt.Println("Connected to Redis")
}

func setString(key, value string) error {
	return rdb.Set(ctx, key, value, 0).Err()
}

func getString(key string) (string, error) {
	return rdb.Get(ctx, key).Result()
}

func main() {
	if err := setString("hello", "world"); err != nil {
		log.Printf("Failed to set string: %v", err)
	}

	val, err := getString("hello")
	if err != nil {
		if err == redis.Nil {
			fmt.Println("Key 'hello' does not exist")
		} else {
			log.Printf("Failed to get string: %v", err)
		}
	} else {
		fmt.Printf("hello: %s\n", val)
	}

	val, err = getString("hello2")
	if err != nil {
		if err == redis.Nil {
			fmt.Println("Key 'hello2' does not exist")
		} else {
			log.Printf("Failed to get string: %v", err)
		}
	} else {
		fmt.Printf("hello2: %s\n", val)
	}
}
