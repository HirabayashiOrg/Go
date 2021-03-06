# 構造体
### 構造体の定義
```
type 構造体名 struct {
  フィールド名 型
  フィールド名 型
  …
}
```
### 構造体の生成
```
s := new(構造体)
```
> `new`では構造体のポインタが返ってくる  

### 値の設定
```
type user struct {
  name  string
  age int
}
=========================
u := new(user)
u.name  = "ryo"
u.score = 25
```
> 構造体.フィールド名でフィールドにアクセスできる  

### 初期化
```
u1 := user{"ryo", 25}
u2 := user{name:"ryo", age:25}
```
> 初期化した場合は構造体自体が返ってくる  
> フィールドの宣言順に{}内に記述することで初期化が可能  
> 丁寧に書くと`field: value`の形式

### 関数の定義
Goでは構造体の内部に関数は持たない  
⇒　構造体の外で定義して構造体と紐づける
```
func (u user) show {
  // 参照するだけの処理
}
func (u *user) update {
  // 更新するような処理
}
```
> 前者は値渡し  
> ⇒　構造体のコピーが渡されるため、参照のみ可能  
> 後者は参照渡し  
> ⇒　値を更新したいようなときはこっち  
