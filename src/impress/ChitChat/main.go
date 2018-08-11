package main

import (
  "fmt"
  "net/http"
)

func main() {
  // マルチプレクサの生成
  mux   := http.NewServeMux()
  // 静的ファイルのパスの指定
  files := http.FileServer(http.Dir("public"))
  // 静的ファイルのハンドリング
  //  リクエストUMLから`/static/`が削除されて静的ファイルを探しにいく
  mux.Handle("/public/", http.StripPrefix("/public/", files))
  // ハンドラの設定
  //  第一引数：URL
  //  第二引数：ハンドラ関数
  mux.HandleFunc("/", index)
  server := &http.Server{
    Addr   : "localhost:6060",
    Handler: mux,
  }
  // サーバ起動
  fmt.Println("server listening @ localhost:6060")
  server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello World !!!")
}
