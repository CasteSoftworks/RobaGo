/*

Scrivere un programma 'pariDispari.go' che legge una sequenza di
valori interi da linea di comando e controlla che si alternino
valori pari e valori dispari.
In questo caso il programma stampa il messaggio "sequenza valida",
altrimenti "elemento in posizione i non valido", dove i è la
posizione del primo  elemento (da sinistra) che non rispetta la
regola di alternanza o che non è un valore numerico.

In caso di mancanza di valori, il programma deve stampare
"nessun valore in ingresso".

La sequenza può iniziare sia con un valore pari sia con uno
dispari.

Si ricorda che lo zero è un numero pari.

Esempi
------

Argomenti:
	3 8 1 12
Output:
	sequenza valida

Argomenti:
	4
Output:
	sequenza valida

Argomenti:
	1 2 3 5
Output:
	elemento in posizione 4 non valido

Argomenti:
	1 2 3eqeqw 5
Output:
	elemento in posizione 3 non valido

Argomenti:

Output:
	nessun valore in ingresso

*/

package main

import(
  "fmt"
  "os"
  "strconv"
)

func controllaVarianza(prec, att int64) (varia bool){
  if (prec%2==0 && att%2!=0)||(prec%2!=0 && att%2==0){
    varia=true
  }
  return
}

func main(){

  if len(os.Args)<2{
    fmt.Println("nessun valore in ingresso")
    os.Exit(1)
  }
  sliceNum:=make([]int64,1)
  primo,err:=strconv.ParseInt(os.Args[1],10,32)
  if err!=nil{
    fmt.Print("elemento in posizione 1 non valido")
    os.Exit(3)
  }
  sliceNum[0]=primo

  for i:=2;i<len(os.Args);i++{
    num,err:=strconv.ParseInt(os.Args[i],10,32)
    if err!=nil{
      fmt.Printf("elemento in posizione %d non valido",i)
      os.Exit(3)
    }

    if !controllaVarianza(sliceNum[i-2],num){
      fmt.Printf("elemento in posizione %d non valido",i)
      os.Exit(2)
    }
    sliceNum=append(sliceNum,num)
  }

  fmt.Print("sequenza valida")
}
