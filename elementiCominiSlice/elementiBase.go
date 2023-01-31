package main

import(
  "fmt"
)

func main(){
  x:=[]string{"ciao","bel","maschione"}
  y:=[]string{"hey","pupa","hai","un","bel","sedere"}

  comp:=0

  for _,e:=range x{
    for _,r:=range y{
      if e==r{
        comp++
      }
    }
  }

  fmt.Println(comp)

}
