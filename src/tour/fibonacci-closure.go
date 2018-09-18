package main

import "fmt"

func fibonacci() func() int {
  sec_last := -1
  last := 1
  return func() int {
    i := last
    last = sec_last + last
    sec_last = i
    return last
  }
}

func main() {
  f := fibonacci()
  for i := 0; i < 10; i++ {
    fmt.Println(f())
  }
}
