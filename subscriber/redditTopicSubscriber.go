package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	subscribe(rdb)
}

func subscribe(rdb *redis.Client) {
	pubsub := rdb.Subscribe(ctx, "reddit_space")
	defer pubsub.Close()

	_, err := pubsub.Receive(ctx)
	if err != nil {
		fmt.Println("Subscribe error:", err)
		return
	}

	ch := pubsub.Channel()
	fmt.Println("Subscribed to reddit_space, waiting for messages...")
	for msg := range ch {
		fmt.Println("Received message:", msg.Payload)
	}
}
