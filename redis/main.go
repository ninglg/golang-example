package main
import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	result, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("ping:", result)
}

func setString() {
	err := rdb.Set(ctx, "hello", "world", 0).Err()
	if err != nil {
		panic(err)
	}
}

func getString() {
	val, err := rdb.Get(ctx, "hello").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("hello:", val)
}

func getNotExistString() {
	val2, err := rdb.Get(ctx, "hello2").Result()
	if err == redis.Nil {
		fmt.Println("hello2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("hello2", val2)
	}
}

func main() {
	setString()
	getString()
	getNotExistString()
}
