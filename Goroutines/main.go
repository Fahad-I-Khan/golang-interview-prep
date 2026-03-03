package main

import (
	"fmt"
	"sync"
)

// Return Unordered output using goroutines
func main() {
	nums := []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup
	for _, num := range nums {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Println(num)
		}(num)
	}
	wg.Wait()
}

