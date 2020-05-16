package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		x := 5
	  fmt.Println(&x)
	}
}
