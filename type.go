package main

import "fmt"

type Integer int 

func (a Integer) Less(b Integer) bool{
  return a<b
}

func main(){
  var a Integer=1
  if(a.Less(2)){
    println(a,"Less 2")
  }

  var aa = [3]int{1, 2, 3}
  var b = aa
  b[1]++
  fmt.Println(aa, b)
}