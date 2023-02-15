/*
Creare una funzione ricorsiva per calcolare una funzione definita così:
per m>0 allora f(n,m) = 1+f(n,m-1)
per m=0 allora f(n,m) = n

aggiunge a n m
*/

package main

import(
  "fmt"
  "os"
  "strconv"
)

func f(n,m int) int{
  if m>0{
    return 1+f(n,(m-1))
  }else{
    return n
  }
}

func main(){
  if len(os.Args)!=3{
    fmt.Println("ERRORE")
    os.Exit(1)
  }

  n,err:=strconv.Atoi(os.Args[1])
  if err!=nil{
    fmt.Println("ERRORE")
    os.Exit(2)
  }
  m,err:=strconv.Atoi(os.Args[2])
  if err!=nil{
    fmt.Println("ERRORE")
    os.Exit(2)
  }

  fmt.Println(f(n,m))
}
