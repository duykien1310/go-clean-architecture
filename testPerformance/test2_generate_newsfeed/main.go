package main

import (
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	numUsers := 125

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	latencies := make([]time.Duration, numUsers)

	start := time.Now()

	for i := 1; i <= numUsers; i++ {
		wg.Add(1)
		go func(userID int) {
			defer wg.Done()
			begin := time.Now()

			err := GenerateNewsfeed(client, userID)
			elapsed := time.Since(begin)
			latencies[userID-1] = elapsed

			if err != nil {
				fmt.Printf("User %d: error -> %v\n", userID, err)
			}
		}(i)
	}

	wg.Wait()
	total := time.Since(start)

	// Tính latency trung bình + p95 + p99
	sort.Slice(latencies, func(i, j int) bool {
		return latencies[i] < latencies[j]
	})

	var totalLatency time.Duration
	for _, l := range latencies {
		totalLatency += l
	}
	avg := totalLatency / time.Duration(numUsers)

	p95 := latencies[int(float64(numUsers)*0.95)-1]
	p99 := latencies[int(float64(numUsers)*0.99)-1]

	fmt.Println("===== Benchmark Result =====")
	fmt.Printf("Tổng thời gian xử lý %d requests: %v\n", numUsers, total)
	fmt.Printf("Average latency: %v\n", avg)
	fmt.Printf("P95 latency: %v\n", p95)
	fmt.Printf("P99 latency: %v\n", p99)
}

func GenerateNewsfeed(client *http.Client, userID int) error {
	url := fmt.Sprintf("http://localhost:8000/api/post/newsfeed/generate?userId=%d", userID)

	req, err := http.NewRequest("GET", url, nil)
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
