package main

import "fmt"

func main() {
  sum := 0
  for i := 0; i < 10; i++ {
    sum += i
  }

  fmt.Println(sum)

  for sum < 100 {
    sum += sum
  }

  fmt.Println(sum)

// infinite loop
  for {

  }

}
