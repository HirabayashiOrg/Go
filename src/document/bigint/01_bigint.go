package main

import (
  "fmt"
  "math/big"
)

func main() {
  // 単純な生成
  n := big.NewInt(1)
  fmt.Println(n)
  // 算術シフト
  n = new(big.Int).Lsh(big.NewInt(5), 3)
  fmt.Println(n)
  //
}
