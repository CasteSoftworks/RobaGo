/*

Scrivere un programma 'squares.go' dotato di:

- una funzione
	func NestedSquares(n int) string
che costruisce e restituisce una stringa corrispondente ad una figura
formata da un quadrato di '*' di dimensione n (con n > 4), che racchiude
un quadrato di '*' di dimensione 'n - 4', separato da quello esterno da una
"cornice" di spazi di spessore unitario (si vedano gli esempi).
La stringa restituita NON deve terminare con '\n'.
Se n < 5, la funzione restituisce la stringa vuota

- una funzione main() che legge da linea di comando un intero
e produce su standard output la figura di quadrati della misura fornita.
Se non ci sono dati sulla linea di comando, il programma stampa il messaggio
"too few arguments" e termina.
Se il valore passato da linea di comando non è un intero, il programma stampa il messaggio
"required integer value" e termina.


Esempi
------
per n = 5
*****
*   *
* * *
*   *
*****

per n = 6
******
*    *
* ** *
* ** *
*    *
******

per n = 7
*******
*     *
* *** *
* * * *
* *** *
*     *
*******

per n = 8
********
*      *
* **** *
* *  * *
* *  * *
* **** *
*      *
********

per n = 9
*********
*       *
* ***** *
* *   * *
* *   * *
* *   * *
* ***** *
*       *
*********
etc..
*/

package main

import(
  "fmt"
  "os"
  "strconv"
)

func NestedSquares(n int) string{
  s:=""

  if n<5{
    return s
  }

  for i:=0;i<n;i++{
    s+="*"
  }
  s+="\n"

  for i:=0;i<n;i++{
    if i==0||i==n-1{
      s+="*"
    }else{
      s+=" "
    }
  }
  s+="\n"

  for i:=0;i<n;i++{
    if i==0||i==n-1{
      s+="*"
    }else{
      if i==1||i==n-2{
        s+=" "
      }else{
        s+="*"
      }
    }
  }
  s+="\n"

  if n>5{
    for j:=0;j<(n-4)-2;j++{
      for i:=0;i<n;i++{
        if i==0||i==n-1{
          s+="*"
        }else{
          if i==1||i==n-2{
            s+=" "
          }else{
            if i==2||i==n-3{
              s+="*"
            }else{
              s+=" "
            }
          }
        }
      }
    s+="\n"
    }

    for i:=0;i<n;i++{
      if i==0||i==n-1{
        s+="*"
      }else{
        if i==1||i==n-2{
          s+=" "
        }else{
          s+="*"
        }
      }
    }
    s+="\n"
  }

  for i:=0;i<n;i++{
    if i==0||i==n-1{
      s+="*"
    }else{
      s+=" "
    }
  }
  s+="\n"

  for i:=0;i<n;i++{
    s+="*"
  }
  s+="\n"

  return s
}

func main(){
  if len(os.Args)!=2{
    fmt.Println("too few arguments")
    os.Exit(1)
  }

  n,err:=strconv.Atoi(os.Args[1])
  if err!=nil{
    fmt.Println("required integer value")
    os.Exit(1)
  }

  fmt.Println(NestedSquares(n))
}
