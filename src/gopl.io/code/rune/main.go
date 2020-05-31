package main

import "fmt"

func main() {
	fmt.Println(string(65))      // "A"
	fmt.Println(string(0x4eac))  // 京
	fmt.Println(string(1234567)) // �
}
