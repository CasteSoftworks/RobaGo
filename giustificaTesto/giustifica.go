package main

import(
  "fmt"
)

func justify(s []string, n int){
  righe:=len(s)
  h:=0

  for i:=0;i<righe;i++{
    lungPa:=len(s[i])

    spazi:=n-lungPa

    fmt.Print("|")

    for j:=0;j<spazi/2;j++{
      fmt.Print(" ")
    }

    fmt.Print(s[i])

    if spazi%2!=0{
      h=1
    }

    for j:=0;j<(spazi+h)/2;j++{
      fmt.Print(" ")
    }

    fmt.Println("|")
  }
}

func main(){

  frase:=[]string{"ciao","mamma","come","stai","?"}
  larghezza:=50

  fmt.Printf("schermo largo: %d\n",larghezza)
  justify(frase,larghezza)
}
