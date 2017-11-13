package main
import (
  "fmt"
  "bufio"
  "os"
)

func main()  {
  fmt.Println("Map相关")
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("输入省份名称:")

  data, _, _ := reader.ReadLine()
  printProvincialCapital(string(data))
}


func printProvincialCapital(provincial string) {
  var provincialCapitalInfo map[string]string
  //不加这句话  将无法在provincialCapitalInfo这个Map里面存放key value
  provincialCapitalInfo = make(map[string]string)
  provincialCapitalInfo["山东"] = "济南"
  provincialCapitalInfo["河北"] = "武汉"
  provincialCapitalInfo["山西"] = "太原"

  provincialCapitalInfo["江苏"] = "南京"

  cityName, ok := provincialCapitalInfo[provincial]
  if ok {
    fmt.Println(provincial, "的省会是", cityName)
  } else {
    fmt.Println("未找到", provincial, "的省会。")

    delete(provincialCapitalInfo, "江苏")

    fmt.Println("目前有如下省份可查：")
    for provincialCapital := range provincialCapitalInfo {
      fmt.Printf("%s省会%s  ", provincialCapital, provincialCapitalInfo[provincialCapital])
    }
    fmt.Println()

  }


}
