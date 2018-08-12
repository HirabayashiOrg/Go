package main

import (
  "crypto/rand"
  "math/big"
  "log"
)

func main() {
  r, err := rand.Int(rand.Reader, big.NewInt(10))
  log.Println(r)
  log.Println(err)
}
