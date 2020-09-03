/*
merge sort algorithm
*/

package main

import (
	"fmt"
)

func main() {
	a := []int{10, 9, 6, 3, 0, 0, -1, -5}
	fmt.Println(mergeSort(a))
}

func mergeSort(a []int) []int {
	aLen := len(a)
	if aLen == 1 {
		return a
	}
	b := mergeSort(a[:aLen/2])
	c := mergeSort(a[aLen/2:])

	d := []int{}
	size, i, j := aLen, 0, 0
	for k := 0; k < size; k++ {
		if i > len(b) && j <= len(c)-1 {
			d = append(d, c[j])
			j++
		} else if j > len(c)-1 && i <= len(b)-1 {
			d = append(d, b[i])
			i++
		} else if b[i] < c[j] {
			d = append(d, b[i])
			i++
		} else {
			d = append(d, c[j])
			j++
		}
	}
	return d
}
