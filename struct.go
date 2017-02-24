package main

import "fmt"

//定义一个struct
type Rect struct{
  x,y float64
  width,height float64
}
//没有构造函数，
func NewRect(x, y, width, height float64) *Rect {
  return &Rect{x, y, width, height}
}

func (r *Rect)Area() float64{
  return r.width * r.height
}

/*---------继承：匿名组合-------------*/
type Base struct{
  Name string
}

/* 类方法 */
func (base *Base) Foo() {
  
}

func (base *Base) Bar(){

}

type Foo struct{
  Base
}

/* 重写Bar方法 */
func (foo *Foo) Bar(){

}

func main(){
  //为什么var obj0 := new(Rect) 会报错：syntax error: unexpected :=
  // 自动推导类型
  obj0 := new(Rect)
  var obj = NewRect(1,2,4,5)
  fmt.Println(obj0.Area(),obj.Area())
}