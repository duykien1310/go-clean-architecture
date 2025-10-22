package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"
)

type PostRequest struct {
	UserId  int    `json:"userId"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	var (
		wg        sync.WaitGroup
		mu        sync.Mutex
		latencies []time.Duration
	)
	numUsers := 125

	start := time.Now()
	client := &http.Client{
		Timeout: 100 * time.Second,
	}

	for i := 1; i <= numUsers; i++ {
		wg.Add(1)
		go func(userID int) {
			defer wg.Done()
			startReq := time.Now()
			err := CreatePost(client, userID)
			duration := time.Since(startReq)

			mu.Lock()
			latencies = append(latencies, duration)
			mu.Unlock()

			if err != nil {
				fmt.Printf("User %d: error -> %v\n", userID, err)
			}
		}(i)
	}

	wg.Wait()
	total := time.Since(start)

	// Sort latencies để tính percentile
	sort.Slice(latencies, func(i, j int) bool {
		return latencies[i] < latencies[j]
	})

	// Tính toán các chỉ số
	avgLatency := average(latencies)
	p95 := percentile(latencies, 95)
	p99 := percentile(latencies, 99)

	fmt.Printf("\n===== Benchmark Report =====\n")
	fmt.Printf("Total users (concurrent): %d\n", numUsers)
	fmt.Printf("Total elapsed time: %v\n", total)
	fmt.Printf("Average latency: %v\n", avgLatency)
	fmt.Printf("P95 latency: %v\n", p95)
	fmt.Printf("P99 latency: %v\n", p99)
	fmt.Println("============================\n")
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

func average(durations []time.Duration) time.Duration {
	var total time.Duration
	for _, d := range durations {
		total += d
	}
	return total / time.Duration(len(durations))
}

func percentile(durations []time.Duration, p float64) time.Duration {
	if len(durations) == 0 {
		return 0
	}
	index := int((p / 100.0) * float64(len(durations)))
	if index >= len(durations) {
		index = len(durations) - 1
	}
	return durations[index]
}
