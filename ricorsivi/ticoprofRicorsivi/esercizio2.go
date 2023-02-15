/*
Creare una funzione ricorsiva che ricevuto un numero restituisce la somma delle cifre del numero se questa è minore di 10 o
il risultato della ri-applicazione della funzione sulla somma delle cifre del numero altrimenti.
Esempi: f(15)=1+5=6, f(392)=f(14)=f(5)=5 dove 3+9+2=14 e 1+4=5.
*/

package main

import(
  "fmt"
  "os"
  "strconv"
)

func f(n int) int{
  res:=0
  nTemp:=n

  for nTemp>0{
    res+=nTemp%10
    nTemp=nTemp/10
  }

  if res<10{
    return res
  }else{
    return f(res)
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
