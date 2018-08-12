# crypto
### 乱数生成
```
r, _ := rand.Int(rand.Reader, big.NewInt(10))
```
> `rand.Int`で乱数を生成  
> ⇒　`Int(rand io.Reader, max *big.Int) (n *big.Int, err error)  `
>   errはエラー発生時に格納される　基本Nil
