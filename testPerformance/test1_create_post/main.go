// Test 1: Fan-out (Create Post)
// 1000 users đồng thời đăng bài
// Mỗi users có 500 followers
// đo tổng thời gian xử lý từ khi gọi CreatePost() đến khi CachePost() cho toàn bộ followers xong.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type PostRequest struct {
	UserId  int    `json:"userId"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	var wg sync.WaitGroup
	numUsers := 1000

	start := time.Now()
	client := &http.Client{
		Timeout: 100 * time.Second,
	}

	for i := 1; i <= numUsers; i++ {
		wg.Add(1)
		go func(userID int) {
			defer wg.Done()
			err := CreatePost(client, userID)
			if err != nil {
				fmt.Printf("User %d: error -> %v\n", userID, err)
			}
		}(i)
	}

	wg.Wait()
	total := time.Since(start)
	fmt.Printf("===> Tổng thời gian xử lý %d CreatePost: %v\n", numUsers, total)
}

func CreatePost(client *http.Client, userID int) error {
	url := "http://localhost:8000/api/post/create"

	reqBody := PostRequest{
		UserId:  userID,
		Title:   fmt.Sprintf("post_user_%d", userID),
		Content: fmt.Sprintf("content_user_%d", userID),
	}

	body, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
