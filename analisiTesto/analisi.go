/*

Analisi di testo (NOTA: CASE SENSITIVE O NO?)
================

Scrivere un programma in Go (il file deve chiamarsi 'analisiTesto.go') dotato di

- un tipo Vocabolario  in cui salvare le stringhe (sequenze di caratteri separate da white spaces) di un testo e, per ciascuna, il numero di volte
che essa appare nel testo

- un metodo String per Vocabolario che stampa il Vocabolario in ordine lessicografico, un lemma per riga seguito da spazio,':', spazio e dalla sua frequenza. Ad esempio:
Lorem : 3
ever : 1

- una funzione

		statistiche(voc Vocabolario) (int, int, float64)

  che restituisce la frequenza più bassa, la frequenza più alta e la media delle frequenze delle stringhe del Vocabolario

- una funzione

		selezione(voc Vocabolario, char rune) []string

	che restituisce la lista delle stringhe contenute nel Vocabolario che iniziano con il carattere (runa) char

- una funzione

		main()

  che legge da linea di comando il nome di un file e un carattere (runa);
  legge il testo contenuto nel file e crea un Vocabolario delle stringhe contenute nel testo, ciascuna con la frequenza con cui appare nel testo;
  stampa il Vocabolario creato, i risultati restituiti dalla funzione statistiche e la lista di stringhe restituita dalla funzione selezione (vedi esempio).
  Se non viene passato un argomento per il nome del file e/o un carattere, il programma deve stampare il messaggio "indicare nome file e iniziale".
  Se si verifica un errore nell'apertura del file, il programma deve stampare l'errore restituito per la mancata apertura del file.
*/
/*
Esempio
-------
(sul file uno.input messo a disposizione)

$ ./analisiTesto uno.input s
Bëatrice : 1
Quant’ : 1
al : 1
bene : 1
che : 2
ch’era : 1
color, : 1
convenia : 1
da : 1
dentro : 1
di : 1
dov’ : 1
entra’mi, : 1
esser : 1
in : 1
io : 1
lucente : 1
lume : 1
l’atto : 1
ma : 1
meglio, : 1
non : 2
parvente! : 1
per : 3
quel : 1
quella : 1
scorge : 1
si : 1
sol : 1
sporge. : 1
subitamente : 1
suo : 1
sé : 1
sì : 2
tempo : 1
È : 1
1 3 1.1388888888888888
[scorge si sol sporge. subitamente suo sé sì]
*/

package main

import(
  "fmt"
  "os"
  "bufio"
  "strings"
  "sort"
)

type Vocabolario struct{
  /*parole []string
  frequenza []int*/
  mappa map[string]int
}

func (d *Vocabolario)String() string{
  var sS []string
  for k,v:=range d.mappa{
    if k==""{
      continue
    }else{
      sS=append(sS,fmt.Sprintf("%s : %d\n",k,v))
    }
  }

  sort.Strings(sS)

  s:=""
  for i:=0;i<=len(sS)-1;i++{
    s+=sS[i]
  }

  return s
}

func statistiche(voc Vocabolario) (int, int, float64){
  var min,max int
  var media float64

  prim:=true
  for k,v:=range voc.mappa{
    if k!=""{
      if prim{
        min=v
        prim=false
      }else{
        if min>v{
          min=v
        }
      }
    }else{
      continue
    }
  }

  prim=true
  for k,v:=range voc.mappa{
    if k!=""{
      if prim{
        max=v
        prim=false
      }else{
        if max<v{
          max=v
        }
      }
    }else{
      continue
    }
  }

  for k,v:=range voc.mappa{
    if k!=""{
      media+=float64(v)
    }else{
      continue
    }
  }
  media/=float64(len(voc.mappa)-1)

  return min,max,media
}

func selezione(voc Vocabolario, char rune) []string{
  var slice []string

  for k,_:=range voc.mappa{
    for _,r:=range k{
      if r==char{
        slice=append(slice,k)
        break
      }else{
        break
      }
    }
  }

  sort.Strings(slice)

  return slice
}

func main(){
  if len(os.Args)!=3{
    fmt.Println("ERRORE")
    os.Exit(1)
  }

  f,err:=os.Open(os.Args[1])
  if err!=nil{
    fmt.Println("ERRORE")
    os.Exit(2)
  }

  var car rune
  for _,r:=range os.Args[2]{
    car=r
  }

  var dizionario Vocabolario
  dizionario.mappa=make(map[string]int)

  var sliceDiz []string

  scanner:=bufio.NewScanner(f)
  for scanner.Scan(){
    linea:=scanner.Text()
    linea=strings.TrimSpace(linea)

    sliceParole:=strings.Split(linea," ")

    for _,p:=range sliceParole{
      sliceDiz=append(sliceDiz,p)
    }
  }

  for _,e:=range sliceDiz{
    _,ok:=dizionario.mappa[e]
    if ok{
      dizionario.mappa[e]+=1
    }else{
      dizionario.mappa[e]=1
    }
  }

  fmt.Println(dizionario.String())
  fmt.Println(statistiche(dizionario))
  fmt.Println(selezione(dizionario,car))
}
