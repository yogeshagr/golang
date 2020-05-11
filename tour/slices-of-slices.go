package main

import (
  "fmt"
  "strings"
  )

func main() {
  // Create a tik tak toe board.
  board := [][]string{
    {"_", "_", "_"},
    {"_", "_", "_"},
    {"_", "_", "_"},
  }

  // The players take turns.
  board[0][1] = "X"
  board[1][0] = "X"
  board[1][2] = "X"
  board[2][1] = "X"

  for i := 0; i < len(board); i++ {
    fmt.Printf("%s\n", strings.Join(board[i], " "))
  }
}
