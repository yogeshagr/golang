package main

import "fmt"

type Vertex struct {
  X int
  Y int
}

func main() {
  fmt.Println(Vertex{2, 3})
  fmt.Println(2)

  v := Vertex{1, 2}
  v.X = 4
  fmt.Println(v.X)

  p := &v
  p.X = 1e9
  fmt.Println(v)

}
