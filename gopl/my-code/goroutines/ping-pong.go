package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	done := make(chan struct{})

	go func() {
		var a string
		for {
			a = <-ch1
			fmt.Println(a)
			ch2 <- "ping"
		}
	}()

	go func() {
		var b string
		for {
			b = <-ch2
			fmt.Println(b)
			ch1 <- "pong"
		}
	}()
	ch2 <- "ping"
	<-done
}
