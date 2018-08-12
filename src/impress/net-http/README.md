# net/httpライブラリを使ったリクエストの受け付け
> ### コラムーカーゴカルト・プログラミングー
> フレークワークの未理解での利用やネットのコードのコピペが該当  
> ⇒　フレームワークであれば約束事やパターンがある以上、  
> 何かしらのトレードオフがあることよ  
> ⇒　ネットのコードはコピペする前にどんなコードなのか理解しようね

### net/httpで使われる構造体
クライアント側：Client, Response  
サーバ側：ResponseWriter, Handler, HandlerFunc, Sercer, ServeMux  
通信：Header, Request, Cookie

### http.ListenAndServe
第一引数：ネットワークアドレス  
第二引数：ハンドラ
> デフォルトは以下の通り  
>   ポート：80
>   ハンドラ：DefaultServerMux

### ハンドラとは？
ハンドラ：ServeHTTPメソッドを持ったインタフェースのこと  
```
ServeHTTP(http.ResponseWriter, *http.Request)
```
### 単純なハンドラ
```
package main

import (
  "fmt"
  "net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "This is my Handler !!!")
}

func main() {
  // ハンドラの生成
  handler := new(MyHandler)
  // サーバインスタンスの生成
  server := http.Server{
    Addr    : "localhost:6060",
    Handler : handler,
  }
  fmt.Println("server listening @ localhost:6060")
  server.ListenAndServe()
}
```
