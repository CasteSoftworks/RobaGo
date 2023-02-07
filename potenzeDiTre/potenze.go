package main

import(
  "fmt"
  "os"
  "strconv"
)

func main(){
  if len(os.Args)<2{
    fmt.Println("ERRORE")
    os.Exit(1)
  }

  for i:=1;i<len(os.Args);i++{
    n,err:=strconv.Atoi(os.Args[i])
    if err!=nil{
      fmt.Println("ERRORE")
      os.Exit(2)
    }

    if n%3==0||n==1{
      if n%2!=0||n==1{
        fmt.Printf("%d ",n)
      }
    }
  }
  fmt.Println()
}
