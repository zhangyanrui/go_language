package main

import (
  "fmt"
)

func main()  {
  fmt.Println("range 学习")

  nums := []int{10, 20, 30, 40}
  for i, num := range nums {
    fmt.Println("index =", i, ", num =", num)
  }

  dictionary := map[string]int {"first" : 0, "second" : 1}
  for k, v := range dictionary {
    fmt.Println("key =", k, ", value =", v)
  }

  for i, c := range "helloworld" {
    fmt.Println("index =", i, ", charCode =", c)
  }
}
