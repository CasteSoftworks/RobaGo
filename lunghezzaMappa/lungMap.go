/*data una mappa con chiave stringa e valore intero,
restituisce in una slice di stringhe le sole chiavi il quale intero ha lo stesso numero di
cifre dei caratteri che compongono la stringa*/

package main

import(
  "fmt"
)

func lung(m map[string]int) (q int){
  for k,v:=range m{
    if len(k)==v{
      q++
    }
  }
  return
}

func main(){
  mappa:=map[string]int{"ciao":4,"cacca":18,"giovanni":3,"cammin":152,"ergastolo":9,}

  fmt.Println(lung(mappa))
}
