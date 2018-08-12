package main

import "net/http"

func main() {
  // デフォルトではポートが80、ハンドラがDefaultServerMuxが使われる
  http.ListenAndServe("", nil)
}
