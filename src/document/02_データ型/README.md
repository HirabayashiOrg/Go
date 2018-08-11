# 基本データ型
### 種類
string  
int  
float64  
bool  
nil：Javaでいうnull  
### 書式付きデータ出力
`fmt.Printf`を用いる  
使用例
```
a := 10
b := "message"
c := true
fmt.Printf("a: %d, b: %s, d: %t", a, b, c)
```
> %d：整数  
> %f：浮動小数  
> %s：文字列  
> %t：真偽値  
> %v：自動判断してくれる  
