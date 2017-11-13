package main
import "fmt"
import "unsafe"
import "math"

// 这种因式分解关键字的写法一般用于声明全局变量(var 或者 const)
var (
		varA int
		varB float64
		varC bool
)

func main (){
	fmt.Println("Go 语言Hello World")

	fmt.Println("Hello, Will is learning go language")
	var age int
	fmt.Println(age)
	var name = "Will"
	fmt.Println(name)
	job := "iOS"
	fmt.Println(job)


	var va, vb, vc = 1, 2, 3
	fmt.Println(va, vb, vc)
	var vd, ve, vf int
	fmt.Println(vd, ve, vf)
	//这种不带声明格式的只能在函数体中出现
	vg, vh, vi := 4, 5, 6
	fmt.Println(vg, vh, vi)
	fmt.Println(varA, varB, varC)

	//这种声明方式声明的变量必须被使用，否则编译器会报错
	aVar := 100
	bVar := 200
	//简单的数据交换 使用并行赋值
	aVar, bVar = bVar, aVar

	fmt.Println(aVar, bVar)

	var cVar int
	//抛弃aVar的值，只接收bVar
	_, cVar = aVar, bVar
	fmt.Println(cVar)


	// 常量可以用len(), cap(), unsafe.Sizeof()常量计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过
	const (
		constA = "abc"
		constB = len(constA)
		constC = unsafe.Sizeof(constA)
	)
	fmt.Println(constA, constB, constC)

	//常量可以作为枚举
	const (
		Unknow = 0
		Female = 1
		Male = 2
	)

	//iota 可以被用作枚举值  在每一个const关键字出现时，被重置为0  然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1
	const (
		Type1 = iota
		Type2 = iota
		Type3 = iota
	)
	fmt.Println(Type1, Type2, Type3)


	// iota的简写
	const (
		Enum1 = iota
		Enum2
		Enum3
	)
	fmt.Println(Enum1, Enum2, Enum3)

	//iota的规律
	const (
            a = iota   //0
            b          //1
            c          //2
            d = "ha"   //独立值，iota += 1
            e          //"ha"   iota += 1
            f = 100    //iota +=1
            g          //100  iota +=1
            h = iota   //7,恢复计数
            i          //8
    )
    fmt.Println(a,b,c,d,e,f,g,h,i)
	//好玩的规律
	const (
		iotaA = 1 << iota
		iotaB
		iotaC = 3 << iota
		iotaD
	)
	fmt.Println(iotaA, iotaB, iotaC, iotaD)

	//不借助第三方变量交换两个数的值
	var number1, number2 = 99, 101
	number1 = number1 ^ number2
	number2 = number2 ^ number1
	number1 = number1 ^ number2
	fmt.Println(number1, number2)

	var ptr *int
	ptr = &number1;
	fmt.Println("ptr = ", ptr)
	fmt.Println("*ptr = ", *ptr)

	var x interface{}
	fmt.Println(x)

	testFunc()
	testClosure()
	testCircle()
}

func testFunc() {
	sqrtFunc := func (r float64) float64 {
		return math.Sqrt(r)
	}
	fmt.Println(sqrtFunc(100))
}



func printC () {
	var c, C  = 1, 0
	A : for C < 100 {
		C++
		for c = 2; c < C; c++ {
			if C % c == 0 {
				goto A
			}
		}
		fmt.Print("素数", C, "\t")
	}
}

func getSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func testClosure() {
	fmt.Println()
	fmt.Println("闭包逻辑")
	getSeq := getSequence()
	fmt.Println(getSeq())
	fmt.Println(getSeq())
	fmt.Println(getSeq())

	getSeq2 := getSequence()
	fmt.Println(getSeq2())
	fmt.Println(getSeq2())
	fmt.Println(getSeq2())

}


type Circle struct {
	radius float64
}

func (c Circle) getArea() float64 {
	return c.radius * c.radius * 3.14
}

func testCircle() {
	var c Circle
	c.radius = 10
	fmt.Println(c.getArea())
}
