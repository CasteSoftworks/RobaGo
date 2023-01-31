package main

import(
  "fmt"
  "os"
)

func main(){
  var numero string

  fmt.Print("metti un numero di 3 cifre, tutte dispari: ")
  fmt.Scan(&numero)

  if len(numero)!=3{
    fmt.Println("\t\tTI HO DETTO DI 3 CIFRE")
    os.Exit(1)
  }

  if int(numero[0])%2==0{
    fmt.Println("\t\tTI HO DETTO CIFRE DISPARI")
    os.Exit(2)
  }
  if int(numero[1])%2==0{
    fmt.Println("\t\tTI HO DETTO CIFRE DISPARI")
    os.Exit(2)
  }
  if int(numero[2])%2==0{
    fmt.Println("\t\tTI HO DETTO CIFRE DISPARI")
    os.Exit(2)
  }

  fmt.Println("Tutte e 3 dispari")
}
