package main

import "fmt"

func main() {
  var a [5]int
  a[2] = 3
  fmt.Println(a)
  // 宣言と初期化を同時に行う
  b := [...]int{1, 2, 3, 4, 5}
  fmt.Println(b)
  fmt.Println(len(b))
}
