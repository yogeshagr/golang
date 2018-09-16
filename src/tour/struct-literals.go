package main

import "fmt"

type Vertex struct {
  X, Y int
}

var (
  v1 = Vertex{1, 2}
  v2 = Vertex{X: 1}
  v3 = Vertex{}
  p1 = &Vertex{1, 2}
  p2 = &v1
)

func main() {
  fmt.Println(v1, v2, v3, p1, &p2, *p1)
}
