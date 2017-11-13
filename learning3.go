package main

import "fmt"

func main()  {
  fmt.Println("Go 语言指针")

  var a int
  fmt.Println("a的地址: ", &a)

  var intB int
  intB = 100
  var iPtr *int
  iPtr = &intB
  fmt.Println("intB内存地址", &intB)
  fmt.Println("iPtr值： ", iPtr)
  fmt.Println("iPtr指向的内存的值： ", *iPtr)

  arr := [3]int{1,2,3}
  fmt.Println("arr =", arr)

  var arrP [3]*int
  fmt.Println("arrP =", arrP)

  goSlice()
}


func goSlice()  {
  fmt.Println("go语言切片")
  a := make([]int, 3)
  fmt.Println("a = ", a)
  var b = []int{1, 2, 3}
  fmt.Println("b = ", b)
  c := b[:]
  fmt.Println("c = ", c)
  d := b[1:2]
  fmt.Println("d = ", d)
  e := b[1:]
  fmt.Println("e = ", e)
  f := b[:1]
  fmt.Println("f = ", f)
  g := make([]int, 10)
  fmt.Println("g = ", g)
  h := make([]int, 2, 10)//第三个参数 用来 指定容量
  fmt.Println("h = ", h)
  fmt.Println("g.cap = ", cap(g), "g.len = ", len(g))
  fmt.Println("h.cap = ", cap(h), "h.len = ", len(h))

  i := make([]int, 0)
  printArray(i)
  j := append(i, 1)
  printArray(j)

}

func printArray(arr []int)  {
   fmt.Println("length:", len(arr), "\tcap:", cap(arr), "\tarr=", arr)
}
