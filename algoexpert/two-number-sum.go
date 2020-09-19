package main

import "sort"

func TwoNumberSum(array []int, target int) []int {
	sort.Ints(array)
	leftPointer := 0
	rightPointer := len(array)-1
	for leftPointer <= rightPointer {
		if array[leftPointer] + array[rightPointer] == target {
			return []int{array[leftPointer], array[rightPointer]}
		} else if array[leftPointer] + array[rightPointer] > target {
			rightPointer--
		} else {
			leftPointer++
		}
	}
	return []int{}
}
