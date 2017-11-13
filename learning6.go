package main
import (
  "fmt"
)

func main()  {
  fmt.Println("接口")

  var phone Phone
  phone = new(IPhone)
  phone.call()

  phone = new(Nokia)
  phone.call()
}

type Phone interface {
  call()
}

type IPhone struct {

}

func (iPhone IPhone) call()  {
  fmt.Println("Call from iPhone")
}

type Nokia struct {

}

func (nokia Nokia) call () {
  fmt.Println("Call from Nokia")
}
