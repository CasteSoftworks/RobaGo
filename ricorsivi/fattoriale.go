package main

import(
  "fmt"
  "strconv"
  "os"
)

func f(n int) int{
  if n==0{
    return 1
  }else{
    return n*f(n-1)
  }
}

func main(){
  if len(os.Args)!=2{
    fmt.Println("Manca argomento numerico")
    os.Exit(1)
  }

  n,err:=strconv.Atoi(os.Args[1])
  if err!=nil{
    fmt.Println("Argomento non numerico")
    os.Exit(2)
  }

  fmt.Println(f(n))
}
