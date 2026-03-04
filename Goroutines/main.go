package main

import (
	"fmt"
	"sync"
)

// Return Unordered output using goroutines
// func main() {
// 	nums := []int{1, 2, 3, 4, 5}
// 	var wg sync.WaitGroup
// 	for _, num := range nums {
// 		wg.Add(1)
// 		go func(num int) {
// 			defer wg.Done()
// 			fmt.Println(num)
// 		}(num)
// 	}
// 	wg.Wait()
// }

// Return Ordered output using goroutines
// type Result struct {
// 	Index int
// 	Value int
// }

// func main() {
// 	nums := []int{1, 2, 3, 4, 5}
// 	ch := make(chan Result, len(nums))
// 	for i, num := range nums {
// 		go func (i, num int)  {
// 			// ch <- Result{i,num}
// 			ch <- Result{Index: i, Value: num}
// 		}(i, num)
// 	}

// 	results := make([]int, len(nums))
// 	for range nums {
// 		res := <- ch
// 		results[res.Index] = res.Value
// 	}
// 	fmt.Println(results)
// }

// Ordered output using worker pool

type Job struct {
	Index int
	Value int
}

type Result struct {
	Index int
	Value int
}

func worker(jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		results <- Result{
			Index: job.Index,
			Value: job.Value,
		}
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	jobs := make(chan Job, len(nums))
	resultch := make(chan Result, len(nums))

	var wg sync.WaitGroup
	workerCount := 3
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(jobs, resultch, &wg)
	}

	// Send jobs to the channel
	go func ()  {
		for i, num := range nums {
			jobs <- Job{Index: i, Value: num}
		}
		close(jobs)
	}()

	// Wait for all workers to finish
	go func ()  {
		wg.Wait()
		close(resultch)
	}()

	results := make([]int, len(nums))

	for res := range resultch {
		results[res.Index] = res.Value
	}
	fmt.Println(results)
}