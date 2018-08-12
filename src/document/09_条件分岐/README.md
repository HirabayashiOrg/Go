# 条件分岐
### if
```
score := 80
if score >= 80 {
  // 処理
} else if score >= 60 {
  // 処理
} else {
  // 処理
}
```
> 条件の記述に()はいらない  

### swich
範囲で分岐させる場合
```
score := 80
switch {
case score >= 80:
  // 処理
case score >= 60:
  // 処理
default:
  // 処理
}
```
値で分岐させる場合
```
color := "red"
switch color {
  case "red":
    // 処理
  default:
    // 処理
}
```
