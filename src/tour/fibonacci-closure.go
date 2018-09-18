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

func fibonacci2() func() int {
    first, second := 0, 1
    return func() int {
        ret := first
        first, second = second, first + second
        return ret
    }
}

func main() {
  f := fibonacci()
  for i := 0; i < 10; i++ {
    fmt.Println(f())
  }

  f2 := fibonacci2()
  for i := 0; i < 10; i++ {
    fmt.Println(f2())
  }

}
