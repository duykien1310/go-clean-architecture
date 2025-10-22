package main

import (
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"
)

func main() {
	numUsers := 125
	var (
		wg        sync.WaitGroup
		mu        sync.Mutex
		latencies []time.Duration
	)

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	start := time.Now()

	for i := 1; i <= numUsers; i++ {
		wg.Add(1)
		go func(userID int) {
			defer wg.Done()

			begin := time.Now()
			err := GenerateNewsfeed(client, userID)
			elapsed := time.Since(begin)

			mu.Lock()
			latencies = append(latencies, elapsed)
			mu.Unlock()

			if err != nil {
				fmt.Printf("User %d: error -> %v\n", userID, err)
			}
		}(i)
	}

	wg.Wait()
	total := time.Since(start)

	// Sort để tính P95, P99
	sort.Slice(latencies, func(i, j int) bool {
		return latencies[i] < latencies[j]
	})

	avg := average(latencies)
	p95 := percentile(latencies, 95)
	p99 := percentile(latencies, 99)
	throughput := float64(numUsers) / total.Seconds()

	fmt.Println("\n===== Benchmark Result =====")
	fmt.Printf("Requests: %d\n", numUsers)
	fmt.Printf("Total time: %v\n", total)
	fmt.Printf("Throughput: %.2f req/s\n", throughput)
	fmt.Printf("Average latency: %v\n", avg)
	fmt.Printf("P95 latency: %v\n", p95)
	fmt.Printf("P99 latency: %v\n", p99)
	fmt.Println("============================\n")
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

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	return nil
}

func average(latencies []time.Duration) time.Duration {
	var total time.Duration
	for _, l := range latencies {
		total += l
	}
	return total / time.Duration(len(latencies))
}

func percentile(latencies []time.Duration, p float64) time.Duration {
	if len(latencies) == 0 {
		return 0
	}
	index := int((p / 100.0) * float64(len(latencies)))
	if index >= len(latencies) {
		index = len(latencies) - 1
	}
	return latencies[index]
}
