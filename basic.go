package main

// import "strings"

import "fmt"

/* 先写参数名后写类型 */
func run(arg string,arg2 int ){
  println(arg)
  println(arg2)
  println("------------")
}

func main() {
  go run("test run",10)
  const LENGTH int = 10
  const WIDTH int = 5   
  var area int
  const a, b, c = 1, false, "str" //多重赋值

  area = LENGTH * WIDTH
  fmt.Printf("面积为 : %d", area)
  println()
  println(a, b, c)   

  var f,l,n=GetName()
  println(f,l,n)


  //枚举
  const (
    Sunday = iota
    Monday
    )
  println(Sunday)
  println(Monday)

  //类型转换
  var value1 int32
  value2 := 64
  value2 = int(value1)

  println(value1,value2)

  // fmt.Println(strings.ToLower("Gopher"))

  //字符串遍历
  str := "Hello,世界"
  var nn = len(str)
  for i := 0; i < nn; i++ {
    ch := str[i] // 依据下标取字符串中的字符，类型为byte
    fmt.Println(i, ch)
  }

  /*---------------数组--------------------*/
  // 先定义一个数组
  var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
  // 基于数组创建一个数组切片
  var mySlice []int = myArray[:5]
  fmt.Println("Elements of myArray: ")
  for _, v := range myArray {
    fmt.Print(v, " ")
  }
  fmt.Println("\nElements of mySlice: ")
  // 数组遍历
  for _, v := range mySlice {
    fmt.Print(v, " ")
  }
  fmt.Println()

  mySlice1 := make([]int, 5)
}

func GetName() (firstName, lastName, nickName string) {
  return "May", "Chan", "Chibi Maruko"
}