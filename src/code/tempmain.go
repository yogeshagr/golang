package main

import (
	"tempconv"
	"fmt"
	)

func main() {
	var c tempconv.Celsius = 32
	fmt.Println(tempconv.CToF(c))

	var k tempconv.Kelvin = 273
	fmt.Println(tempconv.KToC(k))

}
