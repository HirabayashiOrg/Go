# ループ
### 繰り返し数でのループ
```
for i := 0; i < 10; i++ {
  // 処理
}
```
### 終了条件でのループ
javaでいう`while`ループ
```
for {
  // 処理
  if 条件 {
    break
  }
}
```
> Goにはwhileがない  
> forで表現する  

### Rangeを使ったループ
> Javaでいう拡張for文  

配列
```
a := []int{1,2,3,4,5}
for index, value := range a {
  // 処理
}
// 要素番号を使わない場合
for _, value := range a {
  // 処理
}
```
> 配列、スライスの場合は要素番号と要素が渡される  
> 要素番号がいらない場合は`_`で受け取る（破棄する）  

マップ
```
m := map[string]int{"ryo":25, "emi": 24}
for key, value := range m {
  // 処理
}
```
