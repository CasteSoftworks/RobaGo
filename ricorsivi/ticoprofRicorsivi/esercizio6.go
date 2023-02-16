/*
Scrivere una funzione ricorsiva per calcolare l’ennesimo numero della serie di Fibonacci.
La serie di Fibonacci è una serie numerica in cui il numero ennesimo è dato dalla somma dei due numeri di Fibonacci che la precedono;
i primi due numeri di Fibonacci sono 0 e 1.
*/
package main

import(
  "fmt"
  "os"
  "strconv"
)

func f(n int) int{
  if n==0{
    return 0
  }
  if n==1{
    return 1
  }

  return f(n-1)+f(n-2)
}

func main(){
  if len(os.Args)!=2{
    fmt.Println("ERRORE")
    os.Exit(1)
  }

  n,err:=strconv.Atoi(os.Args[1])
  if err!=nil{
    fmt.Println("ERRORE")
    os.Exit(2)
  }

  fmt.Println(f(n))
}
