package main

import (
	"fmt"
  "os"
  "unicode"
)

func controlla(s string) (ok bool,c1,c2,c3,c4,c5 int){
  if len(s)<12{
    c1=1
  }

  for _,l:=range s{
    if unicode.IsSpace(l){
      return
    }
  }

  lettMinM:=0
  lettMaiM:=0
  decM:=0
  altro:=0
  for _,l:=range s{
    if unicode.IsLetter(l){
      if unicode.IsLower(l){
        lettMinM++
      }else{
        lettMaiM++
      }
    }else if unicode.IsDigit(l){
      decM++
    }else{
      altro++
    }
  }

  if lettMaiM>=2&&lettMinM>=2&&decM>=3&&altro>=4{
    ok=true
  }
  if lettMinM<2{
    c2=1
  }
  if lettMaiM<2{
    c3=1
  }
  if decM<3{
    c4=1
  }
  if altro<4{
    c5=1
  }

  return
}

func main(){
  if len(os.Args)!=2{
    fmt.Println("ERRORE")
    os.Exit(1)
  }

  ok,c1,c2,c3,c4,c5:=controlla(os.Args[1])
  if ok{
    fmt.Println("La pw è ben definita!")
  }else{
    fmt.Println("La pw non è definita correttamente:")
    if c1!=0{
      fmt.Println("- La pw deve avere una lunghezza minima di 12 caratteri")
    }
    if c2!=0{
      fmt.Println("- Almeno 2 caratteri della pw devono rappresentare delle lettere minuscole")
    }
    if c3!=0{
      fmt.Println("- Almeno 2 caratteri della pw devono rappresentare delle lettere maiuscole")
    }
    if c4!=0{
      fmt.Println("- Almeno 3 caratteri della pw devono rappresentare delle cifre decimali")
    }
    if c5!=0{
      fmt.Println("- Almeno 4 caratteri della pw non devono rappresentare lettere o cifre decimali")
    }
  }

}
