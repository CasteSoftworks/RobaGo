/**
N.B. leggere il file README.txt per istruzioni di compilazione,
test e consegna

=== Cerchi ===

Scrivere un programma 'cerchi.go' dotato di:

- una struttura 'Cerchio' con campi:
	nome string
	x,y,raggio float64
		(dove x e y sono le coordinate del centro)
  dichiarati in quest'ordine

- una funzione 'newCircle(descr string) Cerchio'
	che data una stringa che descrive un cerchio
	(nome, coordinate x e y del centro, raggio , in quest'ordine e
  separati da spazi)
	restituisce il cerchio corrispondente

- una funzione 'distanzaPunti(x1,y1,x2,y2 float64) float64
	che restituisce la distanza tra i punti di coordinate (x1,y1)
  e (x2,y2)

- una funzione 'tocca(cerc1, cerc2 Cerchio) bool'
	che restituisce true se i due cerchi sono tangenti
	(internamente o esternamente); false altrimenti.
	Trattandosi di valori float, consideriamo una distanza
	trascurabile se è inferiore a 10^-6 (1e-6)

- una funzione maggiore(cerc1, cerc2 Cerchio) bool
	che restituisce true se cerc1 è più grande di cerc2;
	false altrimenti.

- una funzione main()
	che legge da standard input una sequenza (terminata da ctrl D)
	di righe che descrivono cerchi, del tipo:

uno 10.3 12.7 45.8
due 1.3 2.9 5.8
pippo 7.3 22.5 6.6

	ciascuna contenente nome, coordinate del centro e raggio di un
	cerchio, in quest'ordine.

Il programma crea per ciascuna riga il cerchio corrispondente,
lo stampa.
Inoltre stampa se il cerchio, rispetto a quello precedente, è
tangente o no e maggiore o no.
Il primo confronto è fatto rispetto al cerchio "zero"
("", 0, 0, 0).
Non sono richiesti controlli sulla correttezza dei dati in
input, potete assumere che siano sempre nell'ordine e del tipo
previsto.




Esempio
-------
(vengono marcate con '>' le righe digitate da tastiera,
le altre sono l'output del programma)

>uno 0.5 0 2.5
{uno 0.5 0 2.5} non tangente, maggiore
>due 0 0 3
{due 0 0 3} tangente, maggiore
>tre 10.2 -8.4 1.5
{tre 10.2 -8.4 1.5} non tangente, minore o uguale

*/

package main

import(
  "fmt"
  "math"
  "bufio"
  "os"
  "strings"
  "strconv"
)

type Cerchio struct{
  nome string
  x,y,raggio float64
}

func newCircle(descr string) Cerchio{
  var circ Cerchio

  sliceCerch:=strings.Split(descr," ")

  circ.nome=sliceCerch[0]
  circ.x,_=strconv.ParseFloat(sliceCerch[1],64)
  circ.y,_=strconv.ParseFloat(sliceCerch[2],64)
  circ.raggio,_=strconv.ParseFloat(sliceCerch[3],64)

  return circ
}

func distanzaPunti(x1,y1,x2,y2 float64) float64{
  var dist float64

  xMenoX:=x2-x1
  xMenoX=math.Pow(xMenoX,2)

  yMenoY:=y2-y1
  yMenoY=math.Pow(yMenoY,2)

  dist=xMenoX+yMenoY
  dist=math.Sqrt(dist)

  return dist
}

func tocca(cerc1, cerc2 Cerchio) bool{
  var tg bool

  /*
  se cerchi distanzaCentri>somma raggi tg è falsa
  se è minore possono o essere tg o intersecati o interni
    se la distanzaCentri è uguale alla differenza tra i raggi in modulo tg
  */
  sommaRaggi:=cerc1.raggio+cerc2.raggio
  distanzaCentri:=distanzaPunti(cerc1.x,cerc1.y,cerc2.x,cerc2.y)
  diffRagg:=math.Abs(cerc1.raggio-cerc2.raggio)

  //uguali
  if math.Abs(distanzaCentri-sommaRaggi)<0.000001&&math.Abs(distanzaCentri-sommaRaggi)>0{
    tg=true
  }else if distanzaCentri<sommaRaggi{ //uno dentro l'altro
    if math.Abs(diffRagg-distanzaCentri)<0.000001{
      tg=true
    }
  }

  return tg
}

func maggiore(cerc1, cerc2 Cerchio) bool{
  var magg bool

  if cerc2.raggio>cerc1.raggio{
    magg=true
  }

  return magg
}

func main(){
  var prec,att Cerchio
  var primoCiclo bool

  primoCiclo=true
  var zero=Cerchio{
    nome: "",
    x: 0,
    y: 0,
    raggio: 0,
  }

  scanner:=bufio.NewScanner(os.Stdin)
  for scanner.Scan(){
    stringaCerchio:=scanner.Text()

    att=newCircle(stringaCerchio)
    fmt.Print(att)

    if primoCiclo{
      if tocca(zero,att){
        fmt.Print(" tangente")
      }else{
        fmt.Print(" non tangente")
      }

      if maggiore(zero,att){
        fmt.Println(", maggiore")
      }else{
        fmt.Println(", minore o uguale")
      }

      primoCiclo=false
      prec=att

    }else{
      if tocca(prec,att){
        fmt.Print(" tangente")
      }else{
        fmt.Print(" non tangente")
      }

      if maggiore(prec,att){
        fmt.Println(", maggiore")
      }else{
        fmt.Println(", minore o uguale")
      }

      prec=att
    }
  }
}
