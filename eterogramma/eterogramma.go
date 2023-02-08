package main

import (
	"fmt"
  "os"
)

func eterogramma(s string) (ok bool){
  for i,l:=range s{
    for j,h:=range s{
      if i!=j&&l==h&&l!=' '{
        return
      }
    }
  }

  ok=true
  return
}

func main(){
  if len(os.Args)!=2{
    fmt.Println("ERRORE")
    os.Exit(1)
  }

  if eterogramma(os.Args[1]){
    fmt.Println("E' un eterogramma")
  }else{
    fmt.Println("Non è un eterogramma")
  }
}
