package main

import (
  "fmt"
  "net/http"
)

type HelloHandler struct{}
func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello")
}

type ByeHandler struct{}
func (h *ByeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Bye")
}

func main() {
  // ハンドラの生成
  hello  := new(HelloHandler)
  bye    := new(ByeHandler)
  server := http.Server{
    Addr: "localhost:6060",
    // Handler : DefaultServerMux
  }
  http.Handle("/hello", hello)
  http.Handle("/bye"  , bye)

  fmt.Println("server listening @ localhost:6060")
  server.ListenAndServe()
}
