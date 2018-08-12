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
