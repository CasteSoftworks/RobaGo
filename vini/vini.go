/*

Vini
----

Scrivere un programma (il cui file deve chiamarsi 'vini.go') dotato di:

- una struttura BottigliaVino con i seguenti campi (dichiarati nell'ordine):
	nome string
	anno int
	grado float32
	quantita int

- una funzione
	CreaBottiglia(nome string, anno int, grado float32, quantita int) (BottigliaVino, bool)
  che, se i parametri sono **validi** (campi string diversi da vuoto, campi int e float maggiori di zero) crea una bottiglia corrispondente a questi dati e lo restituisce insieme al valore 'true',  altrimenti restituisce una bottiglia "zero-value" e 'false'.

- una funzione main() che crea una slice di bottiglie leggendo da un file (il cui nome è passato da linea di comando) delle righe che contengono ognuna i dati (separati da virgole) relativi ad una bottiglia, ad es:

Rasol,2018,14,750
Camunnorum,2015,15,750
Dom Perignon,2019,12.5,1500
Balon,2013,15,750
Verdicchio,2020,11,375

  e stampa su stdout l'elenco delle bottiglie (esattamente nello stesso formato rappresentato qui sopra).
  Attenzione alle righe vuote (vedere vini.input), il programma deve essere "robusto" e ignorarle.

- una funzione
	CreaBottigliaDaRiga(riga string) (BottigliaVino, bool)
  che crea una bottiglia a partire dalla sua rappresentazione sotto forma di riga di testo, come da formato specificato sopra; se i parametri ci sono tutti e sono validi (vedi sopra), crea una bottiglia corrispondente a questi dati e lo restituisce insieme al valore 'true',  altrimenti restituisce una bottiglia "zero-value" e 'false'.
  Non sono richiesti controlli sui tipi dei dati: si può assumere che i dati, se ci sono, siano nel formato corretto (ma i valori vanno controllati).

- una funzione
	AggiungiBottiglia(bot BottigliaVino, cantina *[]BottigliaVino)
  che aggiunge una bottiglia alla lista

- una funzione
	AggiungiBottigliaDaRiga(bot string, cantina *[]BottigliaVino)
  che effettua la stessa operazione di AggiungiBottiglia ma partendo da una riga di testo (attenzione che è string, nel formato specificato sopra, una riga)

- una funzione
	ContaPerTipo(nome string, cantina []BottigliaVino) int
  che conta quante bottiglie dello stesso tipo (nome) di vino sono presenti nella cantina

- un **metodo** (da applicare a BottigliaVino)
    String() string
  che restituisce una riga di descrizione della bottiglia nel seguente formato:  nome,anno,grado,quantità
  (cioè ad es. restituisca "Rasol,2018,14,750" per la prima riga dell'esempio sopra

*/

package main

import(
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"
)

type BottigliaVino struct{
  nome string
	anno int
	grado float32
	quantita int
}
/*
che, se i parametri sono **validi** (campi string diversi da vuoto, campi int e float maggiori di zero) crea una bottiglia corrispondente a questi
dati e lo restituisce insieme al valore 'true',  altrimenti restituisce una bottiglia "zero-value" e 'false'.
*/
func CreaBottiglia(nome string, anno int, grado float32, quantita int) (BottigliaVino, bool){
  var b BottigliaVino
  var ok bool

  if len(nome)<0{
    return b,ok
  }
  if anno<0{
    return b,ok
  }
  if grado<0{
    return b,ok
  }
  if quantita<0{
    return b,ok
  }

  ok=true

  b.nome=nome
  b.anno=anno
  b.grado=grado
  b.quantita=quantita

  return b,ok
}

/*
una funzione
	CreaBottigliaDaRiga(riga string) (BottigliaVino, bool)
  che crea una bottiglia a partire dalla sua rappresentazione sotto forma di riga di testo, come da formato specificato sopra; se i parametri ci
  sono tutti e sono validi (vedi sopra), crea una bottiglia corrispondente a questi dati e lo restituisce insieme al valore 'true',  altrimenti
  restituisce una bottiglia "zero-value" e 'false'.
  Non sono richiesti controlli sui tipi dei dati: si può assumere che i dati, se ci sono, siano nel formato corretto (ma i valori vanno controllati).
*/
func CreaBottigliaDaRiga(riga string) (BottigliaVino, bool){
  var b BottigliaVino
  var ok bool

  if len(riga)==0{
    return b,ok
  }

  sliceBotSingola:=strings.Split(riga,",")

  a,err:=strconv.Atoi(sliceBotSingola[1])
  if err!=nil{
    return b,ok
  }
  g,err:=strconv.ParseFloat(sliceBotSingola[2],32)
  if err!=nil{
    return b,ok
  }
  q,err:=strconv.Atoi(sliceBotSingola[3])
  if err!=nil{
    return b,ok
  }

  b,ok=CreaBottiglia(sliceBotSingola[0],a,float32(g),q)

  return b,ok
}

/*
una funzione
	AggiungiBottiglia(bot BottigliaVino, cantina *[]BottigliaVino)
  che aggiunge una bottiglia alla lista
*/
func AggiungiBottiglia(bot BottigliaVino, cantina *[]BottigliaVino){
  *cantina=append(*cantina,bot)
}

/*
una funzione
	AggiungiBottigliaDaRiga(bot string, cantina *[]BottigliaVino)
  che effettua la stessa operazione di AggiungiBottiglia ma partendo da una riga di testo (attenzione che è string, nel formato specificato
  sopra, una riga)
*/
func AggiungiBottigliaDaRiga(bot string, cantina *[]BottigliaVino){
  botti,k:=CreaBottigliaDaRiga(bot)
  if k{
    AggiungiBottiglia(botti,cantina)
  }
}

/*
una funzione
	ContaPerTipo(nome string, cantina []BottigliaVino) int
  che conta quante bottiglie dello stesso tipo (nome) di vino sono presenti nella cantina
*/
func ContaPerTipo(nome string, cantina []BottigliaVino) int{
  q:=0
  for _,b:=range cantina{
    if nome==b.nome{
      q++
    }
  }

  return q
}

/*
un **metodo** (da applicare a BottigliaVino)
    String() string
  che restituisce una riga di descrizione della bottiglia nel seguente formato:  nome,anno,grado,quantità
  (cioè ad es. restituisca "Rasol,2018,14,750" per la prima riga dell'esempio sopra
*/

func (b *BottigliaVino)String() string{
  grado:=strconv.FormatFloat(float64(b.grado), 'f', -1, 32)
  return fmt.Sprintf("%s,%d,%s,%d",b.nome,b.anno,grado,b.quantita)
}

/*
- una funzione main() che crea una slice di bottiglie leggendo da un file (il cui nome è passato da linea di comando) delle righe che contengono
ognuna i dati (separati da virgole) relativi ad una bottiglia, ad es:
*/
func main(){
  if len(os.Args)!=2{
    fmt.Println("MANCA NOME FILE")
    os.Exit(1)
  }

  f,err:=os.Open(os.Args[1])
  if err!=nil{
    fmt.Println("ERRORE IN APERURA DI ",os.Args[1])
    os.Exit(2)
  }

  var sliceBottiglie []BottigliaVino
  scanner:=bufio.NewScanner(f)
  for scanner.Scan(){
    AggiungiBottigliaDaRiga(scanner.Text(),&sliceBottiglie)
  }

  for _,b:=range sliceBottiglie{
    fmt.Println(b.String())
  }
}*/
