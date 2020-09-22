package main

import "fmt"

func main() {
	a := []int{-10, 9, 1, -20, 2, 3, 0, 100}
	qSort(a, 0, len(a)-1)
	fmt.Println(a)
}

func qSort(a []int, left, right int) {
	if left >= right {
		return
	}
	pivot := a[(left+right)/2]
	pivotIndex := partition(a, left, right, pivot)

	qSort(a, left, pivotIndex-1)
	qSort(a, pivotIndex, right)
}

func partition(a []int, left, right, pivot int) int {
	for left <= right {
		for a[left] < pivot {
			left++
		}
		for a[right] > pivot {
			right--
		}
		if left <= right {
			a[left], a[right] = a[right], a[left]
			left++
			right--
		}
	}
	fmt.Println(a[:right], left, pivot, a[left])
	return left
}
