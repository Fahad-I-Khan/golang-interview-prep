package main

import "fmt"

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

type Result struct {
	Index int
	Value int
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	ch := make(chan Result, len(nums))
	for i, num := range nums {
		go func (i, num int)  {
			// ch <- Result{i,num}
			ch <- Result{Index: i, Value: num}
		}(i, num)
	}

	results := make([]int, len(nums))
	for range nums {
		res := <- ch
		results[res.Index] = res.Value
	}
	fmt.Println(results)
}