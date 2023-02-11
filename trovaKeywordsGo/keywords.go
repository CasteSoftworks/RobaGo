/*
Esercizio "parole riservate"
===========================

Scrivere un programma Go (il file deve chiamarsi 'goKeywords.go') che identifica le parole riservate
in un programma Go, e per ciascuna stampa le linee di codice in cui si trova.
Se una parola chiave non compare nel testo non viene stampato nulla.
La stampa delle occorrenze delle parole chiave deve seguire l'ordine alfabetico delle stesse.

In particolare il programma deve essere dotato di:

- un funzione printKeywords
che stampa, una per riga e in ordine alfabetico, la lista delle parole riservate trovate, ciascuna seguita dai numeri di linea di codice in cui è stata trovata

- la funzione main
che legge da un file, il cui nome è passato da linea di comando, il testo (un programma Go) da analizzare.
Se manca il parametro da linea di comando, il programma stampa "A file name, please"
Se il file non esiste o non viene aperto correttamente, il programma stampa "File not found".
Altrimenti il programma stampa, una per riga e in ordine alfabetico, la lista delle parole riservate trovate nel file, ciascuna seguita da ':' e dalla lista dei numeri di linea di testo (del programma) in cui è stata trovata.

Elenco delle keywords di Go

"break", "case", "chan", "const", "continue", "default", "defer", "else", "fallthrough", "for",
"func", "go", "goto", "if", "import", "interface", "map", "package", "range", "return", "select",
"struct", "switch", "type", "var"

Esempi
------

$ ./goKeywords hello_world.go.sample
func: [5 9]
import: [3]
package: [1]

dove hello_world.go.sample è il programma che trovate in questa dir.



$ ./goKeywords if.go.sample
else: [9 19 21]
func: [5]
if: [7 13 17 19]
import: [3]
package: [1]

dove if.go.sample è il programma che trovate in questa dir.

*/

package main

import(
  "fmt"
  "os"
  "bufio"
  "strings"
  "sort"
)
/*
un funzione printKeywords
che stampa, una per riga e in ordine alfabetico, la lista delle parole riservate trovate, ciascuna seguita dai numeri di linea di codice
in cui è stata trovata
*/
func printKeywords(m map[string][]int){
  chiavi:=make([]string,0,len(m))

  for k,_:=range m{
    chiavi=append(chiavi,k)
  }
  sort.Strings(chiavi)

  for _,k :=range chiavi{
    fmt.Printf("%s: %v\n",k,m[k])
  }
}

func main(){
  if len(os.Args)!=2{
    fmt.Println("A file name, please")
    os.Exit(1)
  }

  fKey,err:=os.Open("keywords.txt")
  if err!=nil{
    fmt.Println("ERRORE IN APERTURA DI keywords.txt")
    os.Exit(2)
  }

  var mappaKey map[string][]int
  mappaKey=make(map[string][]int)

  scanner:=bufio.NewScanner(fKey)
  for scanner.Scan(){
    mappaKey[scanner.Text()]=nil
  }
  fKey.Close()

  fTesto,err:=os.Open(os.Args[1])
  if err!=nil{
    fmt.Println("File not found")
    os.Exit(2)
  }
  scanner=bufio.NewScanner(fTesto)
  r:=1
  for scanner.Scan(){
    sliceR:=strings.Split(scanner.Text()," ")
    for _,p:=range sliceR{
      _,k:=mappaKey[p]
      if k{
        mappaKey[p]=append(mappaKey[p],r)
      }
    }
    r++
  }
  fKey.Close()

  for k,_:=range mappaKey{
    if mappaKey[k]==nil{
      delete(mappaKey,k)
    }
  }

  printKeywords(mappaKey)
}
