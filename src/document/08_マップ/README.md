# マップ
### 宣言
```
m := make(map[キーの型]値の型)
```
### 値の追加
```
m := make(map[string]int)
m["ryo"] = 25
```
### 宣言＆初期化
```
m := map[string]int{"ryo":25}
m := map[string]int{
  "ryo": 25,
}
```
> 改行を含めて定義をする場合は`,`で終わる  
### ペアの数
```
m := map[string]int{"ryo":25}
len(m)  // 1
```
### 値の取得
```
v, ok := m[key]
```
> マップの値の取得では値と値が存在するかの真偽値が返ってくる  
### 値の削除
```
delete(m, key)
```
> deleteで要素を削除  
>  第一引数：削除対象のマップ  
>  第二引数：削除するのキー  
