# 変数の扱い方
### 変数の宣言
基本パターン
```
var msg string
msg = "message"
```
宣言と代入を同時に
```
var msg = "message"
```
一番簡単
```
msg := "message"
```
### 複数の型を同時宣言
```
var a, b int
var (
  a int
  b string
)
```
### publicとprivate
* 一文字目が大文字：publicで他のパッケージからも参照可能
* 一文字目が小文字：privateで他のパッケージからは参照不可

### 定数
```
const 変数 = 値
```
> 定数宣言の場合は`:=`じゃなくて`=`  
### iota
javaでいうENUM
```
const (
  sun = iota
  mon
  tue
)
```
