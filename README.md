# GoでWebアプリケーション構築
### 環境設定
powershellでの事項を想定
```
cd ※アプリケーションルート
$env:GOPATH=(Convert-Path .)
$env:Path="$env:Path;$env:GOPATH\bin;"
```
> GOはGOPATHを見てビルドを実行するみたい  
> Pathの設定は必須ではないが、bin配下をPathにいれておくと便利  

### コマンド
ビルド
```
go install 作成したアプリケーションのディレクトリ
```

### 用語
マルチプレクサ
> 複数の信号を一つの信号に変換する機構  
> Webアプリケーションで言うとコントローラーのことを指すのかな？  
