/*
Scrivere il codice di una funzione ricorsiva int f(int n) che restituisce quante coppie di cifre uguali in posizioni adiacenti ci sono
nel numero n, nel caso n sia negativo restituisce 0.
Ad es: f(551122) restituisce 3, f(5122) restituisce 1, f(9) restituisce 0, f(444) restituisce 2.
*/

package main

import(
  "fmt"
  "os"
  "strconv"
  "strings"
)

func f(n int) int{
  if n<0{
    return 0
  }else{
    cont:=0
    strN:=""

    for n>0{
      res:=n%10
      strS:=strconv.Itoa(res)
      strN+=strS
      n=n/10
    }

    slS:=strings.Split(strN,"")

    for i:=0;i<len(slS);i++{
      for j:=i+1;j<len(slS);j++{
        if slS[i]==slS[j]{
          cont++
        }
        break
      }
    }
    return cont
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
