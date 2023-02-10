/*
Lessico v2
----------

Copiare il programma 'lessico1.go' in un file 'lessico2.go'

Il programma 'lessico2.go' alimenta un "dizionario" con le parole del testo e le righe (i numeri di riga) in cui ciascuna parola compare; inoltre permette di consultare e stampare il "dizionario".

In particolare, si richiede di dotare il programma 'lessico2.go' delle seguenti funzioni, che andranno usate nel main:

- func PrintMenu()
che stampa il menu delle opzioni, cioè il seguente menu:
+ (legge e memorizza)
? (numeri di riga in cui è comparsa la parola data)
m (stampa le lunghezze min e max)
p (stampa le parole memorizzate)

- func PrintDict(m map[string][]int)
che, data una mappa m, stampa in ordine lessicografico, una per riga, le chiavi della mappa, ciascuna seguita, sulla stessa riga, da uno spazio, un ':', uno spazio e dal suo valore (vedi Esempio sotto)

- AddToDict(line string, n int, dict map[string][]int)
(vedi sotto, opzione '+') che riceve la riga letta che inizia con '+' (line) insieme alla sua posizione nel testo (numero di riga n) e gestisce l'aggiornamento del "dizionario".
Si può assumere, senza fare controlli, che la riga abbia il formato atteso, cioè: '+' seguito da uno spazio e poi da un testo

- Stats(dict map[string][]int) (int, int)
che restiluisce la lunghezza della stringa più corta e la lunghezza della stringa più lunga (in quest'ordine), presenti in dict

Aggiungere inoltre al programma 'lessico2.go', per ogni opzione del menu, dopo le istruzioni di stampa della versione 1, le funzionalità specificate qui sotto.

Quando la riga letta inizia con:
- "+" (alimenta dizionario): il programma usa il rimanente della riga e memorizza in un "dizionario" le parole (stringhe separate da white space) che la costituiscono e il numero di riga. La prima riga letta è la numero 1;
- "?" (consulta dizionario): il programma usa il rimanente della riga come stringa per la consultazione del dizionario e stampa i numeri di riga associati a tale stringa;
- "m" (determina min e max): il programma stampa la lunghezza della stringa più corta e la lunghezza della stringa più lunga (in quest'ordine), presenti nel dizionario
- "p" (print): il programma stampa le parole presenti nel  "dizionario", una per riga e in ordine lessicografico, ciascuna seguita dalla lista dei numeri di riga in cui la parola è comparsa (per il formato, vedi Esempio sotto).


Esempio di esecuzione
---------------------

date in input le seguenti righe di testo:

+ aa bbbb ccc
+ aa
q
p
m
? aa
? aa bbbb

il programma produce il seguente output:

+ (legge e memorizza)
? (numeri di riga in cui è comparsa la parola data)
m (stampa le lunghezze min e max)
p (stampa le parole memorizzate)
alimento il dizionario
alimento il dizionario
opzione non valida
stampo il dizionario ordinato
aa : [1 2]
bbbb : [1]
ccc : [1]
lunghezza min e max
2 4
consulto il dizionario
parola: aa
righe [1 2]
consulto il dizionario
parola: aa bbbb
righe []


*/

package main

import(
  "fmt"
  "os"
  "bufio"
  "strings"
)
/*
che stampa il menu delle opzioni
*/
func PrintMenu(){
  fmt.Println("+ (legge e memorizza)\n? (numeri di riga in cui è comparsa la parola data)\nm (stampa le lunghezze min e max)\np (stampa le parole memorizzate)")
}

/*
che, data una mappa m, stampa in ordine lessicografico, una per riga, le chiavi della mappa, ciascuna seguita, sulla stessa riga, da uno spazio,
un ':', uno spazio e dal suo valore (vedi Esempio sotto)
*/
func PrintDict(m map[string][]int){
 for k,v:=range m{
   fmt.Printf("%s : %v\n",k,v)
 }
}

/*
(vedi sotto, opzione '+') che riceve la riga letta che inizia con '+' (line) insieme alla sua posizione nel testo (numero di riga n) e gestisce
l'aggiornamento del "dizionario". Si può assumere, senza fare controlli, che la riga abbia il formato atteso, cioè: '+' seguito da uno spazio e
poi da un testo
*/
func AddToDict(line string, n int, dict map[string][]int){
  slLinea:=strings.Split(line," ")

  /*var slRiga []int
  slRiga=append(slRiga,n)*/
  for _,p:=range slLinea{
    dict[p]=append(dict[p],n)
  }
}

/*
che restiluisce la lunghezza della stringa più corta e la lunghezza della stringa più lunga (in quest'ordine), presenti in dict
*/
func Stats(dict map[string][]int) (int, int){
  lMin:=0
  for k,_:=range dict{
    if lMin==0{
      lMin=len(k)
    }else{
      if lMin>len(k){
        lMin=len(k)
      }
    }
  }

  lMax:=0
  for k,_:=range dict{
    if lMax==0{
      lMax=len(k)
    }else{
      if lMax<len(k){
        lMax=len(k)
      }
    }
  }

  return lMin,lMax
}

func stampaInfoParola(cerca string,diz map[string][]int){
  fmt.Printf("parola: %s\n",cerca)
  for k,v:=range diz{
    if k==cerca{
      fmt.Printf("righe: %v\n",v)
      return
    }
  }
  fmt.Printf("righe: []\n")
}

func main(){
  var dizionario map[string][]int
  dizionario=make(map[string][]int)

  PrintMenu()

  scanner:=bufio.NewScanner(os.Stdin)
  r:=1
  for scanner.Scan(){
    linea:=scanner.Text()
    //sliceLettere:=strings.Split(linea,"")
    if len(linea)>1{
      if linea[1]==' '{
        if linea[0]=='+'{
          fmt.Println("alimento il dizionario")
          AddToDict(linea[2:],r,dizionario)
        }else if linea[0]=='?'{
          fmt.Println("consulto il dizionario")
          stampaInfoParola(linea[2:],dizionario)
        }else if linea[0]=='m'{
          fmt.Println("lunghezza min e max")
          lMin,lMax:=Stats(dizionario)
          fmt.Println(lMin,lMax)
        }else if linea[0]=='p'{
          fmt.Println("stampo il dizionario ordinato")
          PrintDict(dizionario)
        }else{
          fmt.Println("opzione non valida")
      }
    }
    }else{
      if linea[0]=='+'{
        fmt.Println("alimento il dizionario")
      }else if linea[0]=='?'{
        fmt.Println("consulto il dizionario")
        stampaInfoParola(linea,dizionario)
      }else if linea[0]=='m'{
        fmt.Println("lunghezza min e max")
        lMin,lMax:=Stats(dizionario)
        fmt.Println(lMin,lMax)
      }else if linea[0]=='p'{
        fmt.Println("stampo il dizionario ordinato")
        PrintDict(dizionario)
      }else{
        fmt.Println("opzione non valida")
      }
    }
    r++
  }
}
