// Computes number of bits different between two hashes

package main

import (
	"crypto/sha256"
	"fmt"
)

func bitCount(b byte) int {
	count := 0
	for ; b != 0; count++ {
		b &= b - 1
	}
	return count
}

func shaDiff(a, b []byte) int {
	count := 0
	for i := 0; i < len(a) || i < len(b); i++ {
		switch {
		case i >= len(a):
			count += bitCount(b[i])
		case i >= len(b):
			count += bitCount(a[i])
		default:
			count += bitCount(a[i] ^ b[i])
		}
	}
	return count
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(shaDiff(c1[:], c2[:]))
}
