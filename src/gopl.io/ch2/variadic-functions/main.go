package main

import "fmt"

func max(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	max := vals[0]
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func main() {
	numbers := []int{1, 2, 3, 4, 10, -10, -100}
	fmt.Println(max(numbers...))
	fmt.Println(max())
}
