/*
Scrivere il codice di una funzione ricorsiva f(n) che restituisce:
– n nel caso n sia minore di 10,
– il risultato di f applicata alla somma delle cifre di n se n è pari,
– f(n-1) altrimenti.
*/

package main

import(
  "fmt"
  "os"
  "strconv"
)

func f(n int) int{
  if n<10{
    return n
  }else{
    if n%2==0{
      res:=0
      nTemp:=n

      for nTemp>0{
        res+=nTemp%10
        nTemp=nTemp/10
      }

      return f(res)
    }else{
      return f((n-1))
    }
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
