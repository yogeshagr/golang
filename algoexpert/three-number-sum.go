package main

import (
	"sort"
)

func threeNumberSum(array []int, target int) [][]int {
	triplets := [][]int{}
	sort.Ints(array)
	for i := 0; i < len(array)-2; i++ {
		if array[i] > target {
			break
		}
		leftPtr := i + 1
		rightPtr := len(array) - 1
		for leftPtr < rightPtr {
			currentSum := array[i] + array[leftPtr] + array[rightPtr]
			if currentSum == target {
				triplets = append(triplets, []int{array[i], array[leftPtr], array[rightPtr]})
				leftPtr++
				rightPtr--
			} else if currentSum < target {
				leftPtr++
			} else {
				rightPtr--
			}
		}
	}
	return triplets
}
