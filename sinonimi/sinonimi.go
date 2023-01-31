/*

Sinonimi
========

Si scriva un programma (il file deve chiamarsi 'sinonimi.go') che
gestisca come spiegato qui di seguito l'implementazione di un
dizionario dei sinonimi da applicare ad un testo.

Il programma dovrà lavorare su due file di testo, i cui nomi verranno
passati a riga di comando, il primo conterrà il dizionario dei sinonimi,
il secondo il testo da elaborare (privo di punteggiatura).

Il file del dizionario dei sinonimi conterrà un elenco di parole
corredate del proprio insieme di sinonimi seguendo la seguente
struttura:

	parola: sinonimo1, sinonimo2, sinonimo3, ...

Nota: si può assumere che la struttura delle righe del dizionario dei
sinonimi sia sempre questa senza fare nessun controllo

Il file da elaborare è un semplice testo su più righe, privo di
punteggiatura.
Per semplicità, basta che il testo in output mantenga gli a-capo del
file da elaborare, non occorre invece che mantenga eventuali tab o
spaziature multiple.

Il programma 'sinonimi.go' dovrà accettare come primo parametro il
nome del file dei sinonimi e come secondo parametro il nome del file
da elaborare.
In caso di mancanza di uno dei due parametri il programma dovrà
stampare "2 file names, please".
In caso di assenza di uno dei due file (i cui nomi sono stati passati
a linea di comando) il programma dovrà stampare rispettivamente
"file 1 not found" e "file 2 not found".

L'elaborazione consiste nel produrre un nuovo testo in cui tutte le
parole che possiedono un sinonimo (secondo il file dei sinonimi)
vengono stampate seguite dai loro sinonimi racchiusi tra parentesi
quadre e in ordine alfabetico, ad esempio:

$ ./sinonimi sinonimi.text promessi.input

produce l'output (riportato parzialmente, solo a titolo esemplificativo)

Quel ramo del lago[pozza specchio stagno] di Como che volge a
mezzogiorno tra due catene non interrotte di monti tutto a seni e
a golfi a seconda dello sporgere e del rientrare di quelli vien quasi
a un tratto[pezzo porzione segmento spezzone] a ristringersi e a
prender corso[cammino percorso tragitto] e figura di fiume tra un
promontorio[capo punta] a destra e ampia costiera dall altra parte e
il ponte che ivi congiunge le due rive par che renda ancor più
sensibile all occhio questa trasformazione e segni il punto in cui il
lago[pozza specchio stagno] cessa e l Adda rincomincia per ripigliar
poi nome di lago[pozza specchio stagno] dove le rive allontanandosi di
nuovo lascian l acqua distendersi e rallentarsi in nuovi golfi e in
nuovi seni
...


PROCEDURA 1
mappa con key=parole e value= slice di sinonimi
  copi riga in string
  rimuovi spazi
  slpitti
    salvi parola come key
    rimuovi parola

*/

package main

import(
  "fmt"
  "os"
  "strings"
  "bufio"
  "sort"
)

func main(){
  if len(os.Args)<3{
    fmt.Println("2 file names, please")
    os.Exit(1)
  }

  fileSin,err:=os.Open(os.Args[1])
  if err!=nil{
    fmt.Println("file 1 not found")
    os.Exit(2)
  }

  dizionario:=make(map[string][]string)
  scanner:=bufio.NewScanner(fileSin)
  for scanner.Scan(){
    linea:=scanner.Text()
    linea=strings.ReplaceAll(linea," ","")
    chiave:=strings.Split(linea,":")
    sliceSin:=strings.Split(chiave[1],",")


    dizionario[chiave[0]]=sliceSin
  }
  fileSin.Close()

  fileTesto,err:=os.Open(os.Args[2])
  if err!=nil{
    fmt.Println("file 2 not found")
    os.Exit(2)
  }

  scanner=bufio.NewScanner(fileTesto)
  for scanner.Scan(){
    linea:=scanner.Text()
    lineaSlice:=strings.Split(linea," ")
    for _,parola := range lineaSlice{
      _,ok:=dizionario[parola]
      if ok{
        listaOrd:=dizionario[parola]
        sort.Strings(listaOrd)
        listaSin:=fmt.Sprint(listaOrd)
        linea=strings.ReplaceAll(linea,parola,parola+listaSin)
        for i,s:=range lineaSlice{
          if s==parola{
            lineaSlice[i]=""
          }
        }
      }
    }
    fmt.Println(linea)
  }
  fileTesto.Close()

}
