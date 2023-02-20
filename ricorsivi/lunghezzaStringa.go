package main

import(
  "fmt"
  "os"
)

func contatore(s string) (n int){
  for i,l:=range s{
    if l=='\n'{
      return 0
    }else{
      return 1+contatore(s[i+1:])
    }
  }

  return 0
}

func main(){
  fmt.Println(contatore(os.Args[1]))
}
