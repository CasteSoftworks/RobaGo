/*
Scrivere il codice di una funzione ricorsiva f(n) che restituisce 0 nel caso n sia dispari o zero, 1+f(n/2) altrimenti.
*/
package main

import(
  "fmt"
  "os"
  "strconv"
)

func f(n int) int{
  if n%2!=0||n==0{
    return 0
  }else{
    return 1+f(n/2)
  }
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
