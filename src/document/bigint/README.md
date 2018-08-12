# Bigint
Bigintとは？
* intで扱える範囲を超えた整数を扱いたい場合に使用する

### パッケージ
math/big

### 生成
```
big.NewInt(1)
```

### 算術シフト
左算術シフト
```
new(big.Int).Lsh(対象, シフト数)
new(big.Int).Lsh(big.NewInt(1), 3)   // 1 * 2 ^ 3 = 8
```
> 関数を使用するために`new(big.Int)`で構造体を生成  

右算術シフト
```
new(big.Int).Rsh(big.NewInt(40), 3)  // 40 / ( 2 ^ 3 ) = 5
```
