package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Println("Please provide numbers as argument")
// 		return
// 	}

// 	var numbers []float64
// 	var sum float64

// 	for _, arg := range os.Args[1:] {
// 		if strings.Contains(arg, ",") {
// 			fmt.Println("Invalid Number (contains comma): ", arg)
// 			return
// 		}
// 		if strings.EqualFold(arg, "STOP") {
// 			break
// 		}
// 		n, err := strconv.ParseFloat(arg, 64)
// 		if err != nil {
// 			fmt.Println("Invalid Number: ", arg)
// 			return
// 		}

// 		numbers = append(numbers, n)
// 		sum += n
// 	}
// 	fmt.Println("Numbers:", numbers)
// 	fmt.Println("Sum:", sum)
// }

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var numbers []float64
	var sum float64

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())

		// case in-sensetive
		if strings.EqualFold(text, "STOP") {
			break
		}
		n, err := strconv.ParseFloat(text, 64)
		if err != nil {
			fmt.Println("Invalid number:", text)
			continue
		}

		fmt.Println("Valid number:", n)

		numbers = append(numbers, n)
		sum += n
		no := len(numbers)
		average := sum / float64(no)
		fmt.Printf("Current Sum: %.2f, Average: %.2f\n", sum, average)
	}
	fmt.Println("Numbers:", numbers)
	fmt.Println("Sum:", sum)
}
