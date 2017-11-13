package main
import "fmt"
// import "math"


func main()  {
  fmt.Println("Go 语言数组")
  normalArr()
  forArr()
  multidimensionalArray()
  var strArr = []string{"shan", "dong", "sheng"}
  var strArr2 = [5]string {"xin", "tai", "shi", "xin", "wen"}

  printArr(strArr, 3)
  fmt.Println(strArr2)
}

func normalArr()  {
  fmt.Println()
  fmt.Println("数组语法")
  // Go 声明数组 默认初始化为0
  var intArrA [10] int
  intArrA[3] = 99

  var intArrB = [10]int{2,3,4,5,6,7,8,9,10,11}

  //不确定个数 可以用...
  var intArrC = [...]int{10, 20, 40, 50, 60, 90}

  fmt.Println("intArrA:", intArrA)
  fmt.Println("intArrB:", intArrB)
  fmt.Println("intArrC:", intArrC)
}

func forArr()  {
  fmt.Println()
  fmt.Println("循环语句")

  var floatArr [10] float64
  for i := 0; i < 10; i++ {
    floatArr[i] = 100+float64(i)
  }
  for j := 0; j < 10; j++ {
    fmt.Printf("floatArr[%d]=%.0f\t", j, floatArr[j])
  }
  fmt.Println()
}

func multidimensionalArray()  {
  fmt.Println()
  fmt.Println("多维数组")

  var a = [3][4] int {
    {1,3,5,7},
    {2,4,6,8},
    {9,11,13,15},
  }

  fmt.Println(a)
  fmt.Println("a[2][3] =", a[2][3])
  fmt.Println()
  var b = [5][2] int {
    {1,6},
    {2,7},
    {3,8},
    {4,9},
    {5,10},
  }

  fmt.Println(b)
  for i := 0; i < 5; i++ {
    for j := 0; j < 2; j++ {
      fmt.Printf("b[%d][%d] = %d\t", i, j, b[i][j])
    }
  }
  fmt.Println()
}

func printArr(param[]string, size int)  {
  fmt.Println()
  fmt.Println("函数传递数组")

  for i := 0; i < size; i++ {
    fmt.Println(param[i])
  }
}
