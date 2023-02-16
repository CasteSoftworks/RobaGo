/*
Scrivere una funzione che dato un importo di denaro iniziale, un interesse annuo e un numero di anni permette di calcolare
l’importo di denaro accresciuto degli interessi dopo il numero di anni passati.
*/

package main

import(
  "fmt"
  "os"
  "strconv"
)

func f(anni int,interesse,denaro float64) float64{

  if anni==0{
    return 0
  }else{
    anni--
    return (denaro*interesse*365)/36500.0+f(anni,interesse,denaro)
  }

}

func main(){
  if len(os.Args)!=4{
    fmt.Println("ERRORE")
    os.Exit(1)
  }

  denaro,err:=strconv.ParseFloat(os.Args[1],64)
  if err!=nil{
    fmt.Println("ERRORE")
    os.Exit(2)
  }
  interesse,err:=strconv.ParseFloat(os.Args[2],64)
  if err!=nil{
    fmt.Println("ERRORE")
    os.Exit(2)
  }
  anni,err:=strconv.Atoi(os.Args[3])
  if err!=nil{
    fmt.Println("ERRORE")
    os.Exit(2)
  }

  fmt.Println(denaro+f(anni,interesse,denaro))
}
