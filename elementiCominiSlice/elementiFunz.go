package main

import(
  "fmt"
)

func riempiSliceComune(x,y []string) (ris []string){

  ris=make([]string,1)

  for _,e:=range x{
    for _,r:=range y{
      if e==r{
        ris=append(ris,e)
      }
    }
  }
  ris=ris[1:]

  return
}

func main(){
  x:=[]string{"ciao","bel","maschione"}
  y:=[]string{"hey","pupa","hai","un","bel","sedere"}

  fmt.Println(riempiSliceComune(x,y))

}
