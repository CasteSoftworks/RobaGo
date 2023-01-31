//25/12/1970 era venerdì
package main

import(
  "fmt"
  "os"
  "strconv"
)

type Data struct{
  g,m,a int
}

func dateNatale(sA []int) (sN []Data){
  sN=make([]Data,1)

  for _,a:=range sA{
    var natale Data
    natale.g=25
    natale.m=12
    natale.a=int(a)

    sN=append(sN,natale)
  }
  sN=sN[1:]

  return
}

func main(){
  if len(os.Args)<2{
    fmt.Println("almeno un argomento")
    os.Exit(1)
  }

  var sliceNum []int
  sliceNum=make([]int,1)

  for _,a:=range os.Args[1:]{
    anno,_:=strconv.Atoi(a)
    sliceNum=append(sliceNum,anno)
  }

  sliceNum=sliceNum[1:]

  fmt.Println(dateNatale(sliceNum))

}
