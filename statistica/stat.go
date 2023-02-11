/**

Esercizio statistica
====================

Scrivere un programma in Go (il file deve chiamarsi 'statistica.go') che,
data una serie (di lunghezza arbitraria) di numeri interi da standard input
(la serie contiene almeno un elemento, e termina con CTRL-D su una riga vuota),
calcola ed emette su standard output:

- il valore minimo (int)

- il valore massimo (int)

- la media (in tipo float64, quindi non troncata)

dei valori letti. Per il formato si vedano gli esempi qui sotto.

Si fornisce un file "esempioLungo.input" d'esempio, a fronte del quale l'output *dovrà*
essere il seguente:

min: -3556
max: 9292
media: 2871.39393939394

E un file "esempioBreve.input" d'esempio, a fronte del quale l'output *dovrà*
essere il seguente:

min: -26
max: 77
media: 27.75

*/
package main

import(
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
  "sort"
)

func media(sN []int) (m float64){
  for _,n:=range sN{
    m+=float64(n)
  }

  m/=float64(len(sN))

  return
}

func main(){
  var numeri []int

  scanner:=bufio.NewScanner(os.Stdin)
  for scanner.Scan(){
    sliceS:=strings.Split(scanner.Text()," ")
    for _,n:=range sliceS{
      num,_:=strconv.Atoi(n)
      numeri=append(numeri,num)
    }
  }

  sort.Ints(numeri)
  md:=media(numeri)
  fmt.Printf("min: %d\nmax: %d\nmedia: %f",numeri[0],numeri[len(numeri)-1],md)
}
