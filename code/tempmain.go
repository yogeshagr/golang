package main

import (
	"tempconv"
	"fmt"
	)

func main() {
	var c tempconv.Celsius = 32
	fmt.Println(tempconv.CToF(c))
}
