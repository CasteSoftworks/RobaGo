package main

import(
  "fmt"
  "os"
)

func trovaHost(s string) (host string, ok bool){
  var i, j int

  for i=1;i<len(s);i++{
    if string(s[i])=="/"{
      if string(s[i-1])=="/"{
        for j=i+1;j<len(s);j++{
          if string(s[j])=="/"{
            host=string(s[i+1:j])
            ok=true
            return
          }
        }
      }
    }
  }
  return

}

func main(){

  if len(os.Args)!=2{
    fmt.Println("\t\tManca un indirizzo")
    os.Exit(1)
  }

  indirizzo:=os.Args[1]

  host, ok:=trovaHost(indirizzo)
  if ok==false{
    fmt.Println(indirizzo,"non è un indirizzo valido")
    os.Exit(2)
  }
  fmt.Println(host)
}
