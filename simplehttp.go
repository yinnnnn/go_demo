package main

import(
  // "net"
  "os"
  // "bytes"
  "fmt"
)

// Convert integer to decimal string
func itoa3(val int) string {
  return "hello world"
}


func add(x int, y int) int {
	return x + y
}

func main(){
  fmt.Println(os.Getpid())
  if len(os.Args) !=2{
    fmt.Fprintf(os.Stderr,"Usage: %s host:port",os.Args[0])
    os.Exit(1)
  }
  service := os.Args[1]

  conn,err := net.Dial("tcp",service)
  _,err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))

  result,err := readFully(conn)

  fmt.Println(string(result))

  os.Exit(0)

}
