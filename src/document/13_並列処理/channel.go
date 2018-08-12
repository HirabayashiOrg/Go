package main

import (
  "fmt"
)

func main() {
  myChannel := make(chan int)

  go func() {
    myChannel <- 1
  }()
  
  n := <- myChannel

  fmt.Println(n)
}
