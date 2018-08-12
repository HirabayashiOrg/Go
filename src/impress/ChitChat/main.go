package main

import (
  "fmt"
  "net/http"
)

func main() {
  // マルチプレクサの生成
  mux   := http.NewServeMux()
  // 静的ファイルのパスの指定
  files_js   := http.FileServer(http.Dir("public/js"))
  files_css  := http.FileServer(http.Dir("public/css"))
  files_html := http.FileServer(http.Dir("public/html"))
  // 静的ファイルのハンドリング
  //  リクエストUMLから`/static/`が削除されて静的ファイルを探しにいく
  mux.Handle("/js/"  , http.StripPrefix("/js/" , files_js))
  mux.Handle("/css/" , http.StripPrefix("/css/", files_css))
  mux.Handle("/html/", http.StripPrefix("/html/", files_html))
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
