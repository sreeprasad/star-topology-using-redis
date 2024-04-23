package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	go func() {
		for {
			if err := fetchAndPublish(rdb); err != nil {
				fmt.Println("Error:", err)
			}
			time.Sleep(10 * time.Second)
		}
	}()

	select {}
}

func fetchAndPublish(rdb *redis.Client) error {
	url := "https://www.reddit.com/r/space/top.json?limit=10"
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("User-Agent", "go:redditbot:v0.1 (by /u/sreeprasad)")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	posts := result["data"].(map[string]interface{})["children"].([]interface{})

	for _, post := range posts {
		postData := post.(map[string]interface{})["data"].(map[string]interface{})
		title := postData["title"].(string)

		if err := rdb.Publish(ctx, "reddit_space", title).Err(); err != nil {
			return err
		}
	}

	fmt.Println("Published latest post titles to Redis.")
	return nil
}
