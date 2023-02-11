/*
# TESTO

Realizzare un programma Go (nome file 'uniq.go') che implementi un semplice 'uniq', comando Unix che trova le linee consecutive ripetute in un file
di testo e produce una specie di compressione del file indicando quante volte è ripetuta ogni riga.

Il programma deve prendere un parametro da linea di comando:
- nome del file da elaborare

e stampare su standard output le linee del file, senza stampare le ripetizioni consecutive, con ciascuna linea preceduta dal numero di volte che la
linea è ripetuta nel file. Separare la linea dal numero con uno spazio bianco.

Nota bene: NON va implementato invocando da Go il comando 'uniq' di sistema.

## Esempio esecuzione

(presuppone il 'uniq.go' già compilato, 'uno.input' è presente in questa directory)

$ ./uniq uno.input
1 Lorem Ipsum is simply dummy text of the printing
1 and typesetting industry. Lorem Ipsum has been the industry's standard dummy
1 text ever since the 1500s, when an unknown printer took
1 a galley of type and scrambled it to make a type
1 specimen book. It has survived not only five centuries, but also the leap into electronic typesetting,
1 remaining essentially unchanged. It was popularised in the 1960s with the release
1 of Letraset sheets containing Lorem Ipsum passages, and more
1 recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.
2 Lorem Ipsum is simply dummy text of the printing
1 and typesetting industry. Lorem Ipsum has been the industry's standard dummy
1 text ever since the 1500s, when an unknown printer took
4 a galley of type and scrambled it to make a type
1 specimen book. It has survived not only five centuries, but also the leap into electronic typesetting,
1 remaining essentially unchanged. It was popularised in the 1960s with the release
1 of Letraset sheets containing Lorem Ipsum passages, and more
1 recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.

*/

package main

import(
  "fmt"
  "os"
  "bufio"
)

func main(){
  if len(os.Args)!=2{
    fmt.Println("errore")
    os.Exit(1)
  }

  f,err:=os.Open(os.Args[1])
  if err!=nil{
    fmt.Println("errore")
    os.Exit(2)
  }

  /*var mappaFrasi map[string]int
  mappaFrasi=make(map[string]int)*/
  var sliceTesto []string

  scanner:=bufio.NewScanner(f)
  for scanner.Scan(){
    lE:=scanner.Text()
    sliceTesto=append(sliceTesto,lE)
  }


  for i:=0;i<len(sliceTesto);i++{
    q:=1
    for j:=i+1;j<len(sliceTesto);j++{
      if sliceTesto[i]==sliceTesto[j]{
        q+=1
      }else{
        fmt.Println(q,sliceTesto[i])
        break
      }
    }
    if q!=1{
      i+=q-1
    }
  }
  fmt.Println("1",sliceTesto[len(sliceTesto)-1])
}
