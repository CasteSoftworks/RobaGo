package main

import (
	"fmt"
)

func controlla(s string) (ok bool, q int){
  var vocale rune

  prim:=true
  for _,e:=range s{
    if e=='a'||e=='e'||e=='i'||e=='o'||e=='u'{
      if prim{
        q++
        vocale=e
        prim=false
      }else{
        if e==vocale{
          q++
        }else{
          return
        }
      }
    }
  }
  ok=true

  return
}

func main(){
  var s string

  for {
    fmt.Println("inserisci una parola")
    fmt.Scan(&s)
    if s=="0"{
      break
    }
    k,n:=controlla(s)
    if k{
      fmt.Println("corretto a",n)
    }else{
      fmt.Println("sbagliato")
    }
  }
}
