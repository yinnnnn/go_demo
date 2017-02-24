package main
import "fmt"
import "sync"
import "runtime"
var counter int = 0

//c++实现版本
func Count(lock *sync.Mutex) {
  lock.Lock()
  counter++
  fmt.Println(counter)
  lock.Unlock()
}

func Count_go(ch chan int,i int) {
  if i!=2{
    ch <- i  
  }
  
  fmt.Println("Counting ")
}
func main() {
  lock := &sync.Mutex{}
  for i := 0; i < 10; i++ {
    go Count(lock)
  }
  for {
    lock.Lock()
    c := counter
    lock.Unlock()
    runtime.Gosched()
    if c >= 10 {
      break
    }
  }

  chs := make([]chan int,10)
  for i:=0; i<10; i++ {
    chs[i] = make(chan int)
    go Count_go(chs[i],i)
  }

  for _,ch := range(chs){
    var aa=<-ch
    fmt.Println(aa,"-------")
  }
}