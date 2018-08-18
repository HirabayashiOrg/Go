package main

import (
  "fmt"
  "net/http"
)

type HelloHandler struct{}
func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello")
}

func bye(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Bye")
}

func main() {
  server := http.Server{
    Addr: "localhost:6060",
    // Handler : DefaultServerMux
  }
  // ハンドラの利用
  hello  := new(HelloHandler)
  http.Handle("/hello", hello)
  // ハンドラ関数利用
  http.HandleFunc("/bye"  , bye)

  fmt.Println("server listening @ localhost:6060")
  server.ListenAndServe()
}
