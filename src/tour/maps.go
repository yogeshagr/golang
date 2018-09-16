package main

import "fmt"

type Vertex struct{
  Lat, Long float64
}

var m map[string]Vertex

func main() {
  m = make(map[string]Vertex)
  m["Bell labs"] = Vertex{
    40.68433, -74.39967,
  }

  m["Google"] = Vertex{2, 3}

  var n = map[string]Vertex {
    "Bell labs" : {
      40.68, -74.2,
    },
    "Google" : {
      37.7, -122.2,
    },
  }

  fmt.Println(m["Bell labs"])
  fmt.Println(m)
  fmt.Println(n)
}
