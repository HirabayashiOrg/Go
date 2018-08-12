package main

import "fmt"

type user struct {
  name string
  age  int
}

func (u user) show() {
  fmt.Printf("name: %s, age: %d\n", u.name, u.age)
}

func (u *user) glow() {
  u.age++
}

func main() {
  // newで構造体を生成
  pu := new(user)
  // u には構造体のポインタが格納されている
  fmt.Println(pu)
  // 値の格納
  pu.name = "ryo"
  pu.age  = 25
  fmt.Println(pu)
  // 初期化 ここでは構造体自体が格納される
  u := user{
    name: "ryo",
    age : 25,
  }
  fmt.Println(u)

  // ageを更新
  u.glow()
  // 表示
  u.show()
}
