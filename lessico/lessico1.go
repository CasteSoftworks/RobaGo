/*
Lessico v1
----------
Scrivere un programma 'lessico1.go' che si comporta come segue.

1) Stampa il seguente menu di opzioni:
  + (legge e memorizza)
  ? (numeri di riga in cui è comparsa la parola data)
  m (stampa le lunghezze min e max)
  p (stampa le parole memorizzate)

2) Legge da standard input un testo di un numero arbitrario di righe e termina quando riceve un "end of file" (cioè EOF, da tastiera pressione di
'CTRL-d' dopo aver dato <invio>).

Quando la riga letta **inizia** con:
- "+" : il programma stampa "alimento il dizionario"
- "?" : il programma stampa "consulto il dizionario"
- "m" : il programma stampa "lunghezza min e max"
- "p" : il programma stampa "stampo il dizionario ordinato"
- per qualsiasi altro carattere il programma stampa "opzione non valida"

Esempio di esecuzione
---------------------

date in input le seguenti righe di testo:

  + la befana ha il fazzoletto e la gonna rattoppata
  ? ha
  n
  + ma quest'anno poverina
  + la befana è raffreddata!
  ? ha il
  m
  p

il programma produce il seguente output:

  + (legge e memorizza)
  ? (numeri di riga in cui è comparsa la parola data)
  m (stampa le lunghezze min e max)
  p (stampa le parole memorizzate)
  alimento il dizionario
  consulto il dizionario
  opzione non valida
  alimento il dizionario
  alimento il dizionario
  consulto il dizionario
  lunghezza min e max
  stampo il dizionario ordinato

*/

package main

import(
  "fmt"
  "os"
  "bufio"
  "strings"
)

func main(){
  fmt.Println("+ (legge e memorizza)\n? (numeri di riga in cui è comparsa la parola data)\nm (stampa le lunghezze min e max)\np (stampa le parole memorizzate)")

  scanner:=bufio.NewScanner(os.Stdin)
  for scanner.Scan(){
    linea:=scanner.Text()
    sliceLettere:=strings.Split(linea,"")
    if len(sliceLettere)>1{
      if sliceLettere[1]==" "{
        if sliceLettere[0]=="+"{
          fmt.Println("alimento il dizionario")
        }else if sliceLettere[0]=="?"{
          fmt.Println("consulto il dizionario")
        }else if sliceLettere[0]=="m"{
          fmt.Println("lunghezza min e max")
        }else if sliceLettere[0]=="p"{
          fmt.Println("stampo il dizionario ordinato")
        }else{
          fmt.Println("opzione non valida")
      }
    }
    }else{
      if sliceLettere[0]=="+"{
        fmt.Println("alimento il dizionario")
      }else if sliceLettere[0]=="?"{
        fmt.Println("consulto il dizionario")
      }else if sliceLettere[0]=="m"{
        fmt.Println("lunghezza min e max")
      }else if sliceLettere[0]=="p"{
        fmt.Println("stampo il dizionario ordinato")
      }else{
        fmt.Println("opzione non valida")
      }
    }
  }
}
