package main

import(
  "fmt"
  "os"
)

func contatore(s string, c rune) (n int){
  for i,l:=range s{
    if l==c{
      return 1+contatore(s[i+1:],c)
    }
  }

  return 0
}

func main(){

  var c rune
  for _,l:=range os.Args[2]{
    c=l
    break
  }

  fmt.Println(contatore(os.Args[1],c))
}
