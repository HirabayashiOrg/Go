package main

import "fmt"

func main() {
  m := make(map[string]int)
  m["ryo"] = 25
  m["emi"] = 24
  fmt.Println(m)

  m = map[string]int{"ryo" : 25,"emi" : 24}
  fmt.Println(m)
  fmt.Println(m["ryo"])
  v, ok := m["ryo"]
  fmt.Printf("v: %v, ok: %t\n", v, ok)

  m = map[string]int{
    "ryo": 25,
  }
  fmt.Println(m)
}
