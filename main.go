package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}

	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(size) + 1
	}

	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for _, value := range data[1:] {
		if value > max {
			max = value
		}
	}

	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	var wg sync.WaitGroup
	chunkMaximums := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS

	for i := range CHUNKS {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)
		}

		wg.Add(1)
		go func(chunk []int, index int) {
			defer wg.Done()
			chunkMaximums[index] = maximum(chunk)
		}(data[start:end], i)
	}

	wg.Wait()

	return maximum(chunkMaximums)
}

func main() {
	fmt.Printf("Generating %d integers\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Looking for the maximum value in a single thread")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Maximum element value: %d\nSearch time: %d mcs\n", max, elapsed)

	fmt.Printf("Looking for the maximum value in %d threads\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Maximum element value: %d\nSearch time: %d mcs\n", max, elapsed)
}
