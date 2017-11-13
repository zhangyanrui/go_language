package main
import (
  "fmt"
  "math"
  "errors"
)

func main()  {
  fmt.Println("Error的使用")
  result, err := sqrt(-1)

  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(result)
  }

  resultInt, errorStr := Divide(1,0)
  if errorStr != ""{
    fmt.Println(errorStr)
  } else {
    fmt.Println(resultInt)
  }
}

func sqrt(f float64) (float64, error) {
  if f<0 {
    return 0, errors.New("math: 不能对负数求平方根")
  }
  return math.Sqrt(f), nil
}


type DivideError struct {
  dividee int
  divider int
}

func (deErr *DivideError)Error() string {
  strFormat := `
  不能执行，除数不能为0.
    被除数: %d
    除数: 0
  `
  return fmt.Sprintf(strFormat, deErr.dividee)
}

func Divide(varDividee int, varDivider int) (int, string) {
  if varDivider == 0 {
    dError := DivideError {
      dividee: varDividee,
      divider: varDivider,
    }
    return 0, dError.Error()
  } else {
    return varDividee / varDivider , ""
  }
}
