package main

import "fmt"

func main() {
  m := map[string]int{"ryo":25, "emi":24}
  for k, v := range m {
    fmt.Printf("k: %s, v: %d\n", k, v)
  }
}
