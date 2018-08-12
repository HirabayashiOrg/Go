package main

import "fmt"

func main() {
  a := []int{1,2,3,4,5}
  b := a[:]
  fmt.Println(b)
  c := a[:3]
  fmt.Printf("len: %d, cap: %d\n", len(c), cap(c))

  s := make([]int, len(a))
  n := copy(s, a)
  fmt.Printf("s: %v, n: %d\n", s, n)
}
