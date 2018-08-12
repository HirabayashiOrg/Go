package main

import "fmt"

func task1(myChan chan string) {
  for i:=0; i<100; i++ {
    fmt.Println("doing task1: %d", i)
  }
  myChan <- "task1 finished!"
}

func task2(myChan chan string) {
  for i:=0; i<100; i++ {
    fmt.Println("doing task2: %d", i)
  }
  myChan <- "task2 finished!"
}

func main() {
  myChan1 := make(chan string)
  myChan2 := make(chan string)

  go task1(myChan1)
  go task2(myChan2)

  result1 := <- myChan1
  result2 := <- myChan2

  fmt.Println(result1)
  fmt.Println(result2)
}
