package main

import (
  "fmt"
  "math"
)

type I interface {
  M()
}

func (t *T) M() {
  fmt.Println(t.S)
}

func (f F) M() {
  fmt.Println(f)
}

type T struct {
  S string
}

type F float64

func main(){
  var i I

  i = &T{"Hello"}
  describe(i)
  i.M()

  i = F(math.Pi)
  describe(i)
  i.M()

}

func describe(i I){
  fmt.Printf("(%v, %T)\n", i, i)
}
