package main

import "fmt"

func plus(a int, b int) int {
  return a + b
}

func plusMinus(a int, b int) (int, int) {
  return a + b, a - b
}

func io(a int) int {
  return a
}

func main() {
  a := 1
  b := 2
  fmt.Println(plus(a, b))
  var c, d = plusMinus(a, b)
  fmt.Printf("plus: %d, minus: %d\n", c, d)
  fmt.Println(io(1))
}
