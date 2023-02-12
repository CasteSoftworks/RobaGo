/*
Scrivere un programma 'distanzaNumerica.go' che legge da linea di comando
una serie di numeri interi e verifica che la distanza numerica (differenza in valore assoluto)
fra due elementi consecutivi della sequenza sia sempre corrispondente
al numero di argomenti passati.
Se lo è, stampa la somma dei numeri; altrimenti stampa NON OK.
Anche nel caso la sequenza contenga nessun o un solo elemento, il programma stampa NON OK.

Esempi
------

$ go run distanzaNumerica.go 1 4 7
12

$ go run distanzaNumerica.go 10 7 7 7 8
NON OK

$ go run distanzaNumerica.go 10 5 0 -5 0
10

*/

package main

import(
  "fmt"
  "os"
  "strconv"
  "math"
)

func main(){
  if len(os.Args)<2{
    fmt.Println("NON OK")
    os.Exit(1)
  }

  var numeri []int
  for i:=1;i<len(os.Args);i++{
    n,err:=strconv.Atoi(os.Args[i])
    if err!=nil{
      fmt.Println("NON OK")
      continue
    }

    numeri=append(numeri,n)
  }

  diffIni:=math.Abs(float64(numeri[1]-numeri[0]))
  somma:=numeri[0]+numeri[1]
  
  for i:=2;i<len(numeri);i++{
    diffAtt:=math.Abs(float64(numeri[i]-numeri[i-1]))
    if diffAtt!=diffIni{
      fmt.Println("NON OK")
      os.Exit(1)
    }else{
      somma+=numeri[i]
    }
  }

  fmt.Println(somma)
}
