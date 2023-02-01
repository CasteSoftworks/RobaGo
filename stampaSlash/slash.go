package main

import(
  "fmt"
)

func slash(h int){
  for i:=0;i<h;i++{
    for j:=0;j<h-1-i;j++{
      fmt.Print(" ")
    }
    fmt.Println("*")
  }
}

func backSlash(h int){
  for i:=0;i<h;i++{
    for j:=0;j<i;j++{
      fmt.Print(" ")
    }
    fmt.Println("*")
  }
}

func main(){
  slash(6)
  backSlash(6)
}
