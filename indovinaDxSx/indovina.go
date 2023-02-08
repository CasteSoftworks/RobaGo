package main

import (
	"fmt"
  "math/rand"
  "time"
)

func main(){
  var vtt, prs int
  var r string


  rand.Seed(time.Now().UnixNano())

  for{
    dir:=rand.Intn(2)
    fmt.Println("Dove sto guardando?")
    fmt.Scan(&r)
    if r=="s"{
      if dir==0{
        fmt.Println("Giusto, hai vinto!")
        vtt++
        continue
      }
    }else if r=="d"{
      if dir==1{
        fmt.Println("Giusto, hai vinto!")
        vtt++
        continue
      }
    }else if r=="0"{
      break
    }
    fmt.Println("Sbagliato, hai perso!")
    prs++
  }

  tot:=prs+vtt
  fmt.Printf("Vinto: %d%% Perso: %d%%\ns",(vtt*100)/tot,(prs*100)/tot)

}
