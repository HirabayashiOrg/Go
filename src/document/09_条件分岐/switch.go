package main

import "fmt"

func main() {
  score := 80
  if score >= 80{
    fmt.Println("Great!!")
  } else if score >= 60 {
    fmt.Println("Good!")
  } else {
    fmt.Println("bad score...")
  }
}
