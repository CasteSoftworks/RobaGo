/*
draw
------

Scrivere un programma 'draw.go' dotato di:

- una funzione
	main()
che legge come parametri da linea di comando (in quest'ordine)
due numeri interi l e n e un carattere c (byte),
e, se l e n sono > 0, produce su standard output una scala di n gradini
di lunghezza e altezza l disegnati usando il carattere c (vedi esempi sotto),
altrimenti non fa niente.

Si può assumere che il programma venga lanciato con tre parametri
dei tipi attesi (non occorre cioè fare controlli).


Esempi
------

$ go run draw.go 3 1 x
xxx
  x
  x



$ go run draw.go 3 2 +
+++
  +
  +
  +++
    +
    +

*/
package main

import(
  "fmt"
  "os"
  "strconv"
)
/*
una funzione
	DrawPoint(c byte, k int) string
che restituisce una stringa formata da k spazi bianchi
seguiti dal carattere c*/
func DrawPoint(c byte, k int) string{
  var s string

  for i:=0;i<k;i++{
    s+=" "
  }
  s+=string(c)

  return s
}

/*
una funzione
	DrawSegment(c byte, k, l int) string
che restituisce una stringa formata da k spazi bianchi
seguiti da l caratteri c*/
func DrawSegment(c byte, k, l int) string{
  var s string

  for i:=0;i<k;i++{
    s+=" "
  }

  for i:=0;i<l;i++{
    s+=string(c)
  }

  return s
}

/*
una funzione
	main()
che legge come parametri da linea di comando (in quest'ordine)
due numeri interi l e n e un carattere c (byte),
e, se l e n sono > 0, produce su standard output una scala di n gradini
di lunghezza e altezza l disegnati usando il carattere c (vedi esempi sotto),
altrimenti non fa niente.
*/
func main(){
  if len(os.Args)!=4{
    fmt.Println("ERRORE")
    os.Exit(1)
  }

  l,err:=strconv.Atoi(os.Args[1])
  if err!=nil||l<=0{
    fmt.Println("ERRORE")
    os.Exit(2)
  }
  n,err:=strconv.Atoi(os.Args[2])
  if err!=nil||n<=0{
    fmt.Println("ERRORE")
    os.Exit(2)
  }
  if len(os.Args[3])!=1{
    fmt.Println("ERRORE")
    os.Exit(2)
  }
  c:=os.Args[3][0]

  spazioS:=l-1
  spazioP:=spazioS
  for i:=0;i<n;i++{
    fmt.Println(DrawSegment(c,i*spazioS,l))
    for j:=0;j<spazioS;j++{
      fmt.Println(DrawPoint(c,spazioP))
    }
    spazioP+=spazioS
  }
}
