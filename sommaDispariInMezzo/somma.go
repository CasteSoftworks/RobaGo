package main

import (
	"fmt"
  "os"
  "strconv"
)

func main(){
  if len(os.Args)!=3{
    fmt.Println("ERRORE")
    os.Exit(1)
  }

  n1,err:=strconv.Atoi(os.Args[1])
  if err!=nil{
    fmt.Println("ERRORE")
    os.Exit(2)
  }
  n2,err:=strconv.Atoi(os.Args[2])
  if err!=nil{
    fmt.Println("ERRORE")
    os.Exit(2)
  }

  sommaDisp:=0
  if n1<n2{
    for i:=n1+1;i<n2;i++{
      if i%2!=0{
        sommaDisp+=i
      }
    }
  }else if n1>n2{
    for i:=n2+1;i<n1;i++{
      if i%2!=0{
        sommaDisp+=i
      }
    }
  }

  fmt.Println(sommaDisp)

}
