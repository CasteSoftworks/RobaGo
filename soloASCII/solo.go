package main

import(
  "fmt"
)

func soloASCII(s string) (r string){
  r+="\n"
  for _,e:=range s{
    if len(string(e))==1{
      r+=string(e)
    }
  }

  return r
}

func main(){
  s:="⌘questa ée' una stringa 語 formatàa di 本 soli ASCII⌘\n"
  fmt.Print(soloASCII(s))
}
