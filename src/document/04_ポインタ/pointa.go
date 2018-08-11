package main

import "fmt"

func main() {
  // 普通の変数宣言
  a := 1
  // ポインタ用の変数宣言
  var pa *int
  // aのアドレスを代入
  pa = &a

  fmt.Println(a)
  fmt.Println(pa)
  fmt.Println(*pa)
}
