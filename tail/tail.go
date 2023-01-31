/*

Realizzare un programma Go (nome file 'tail.go') che implementi
un semplice 'tail', comando Unix che estrae le ultime N righe
di un file di testo.

Il programma deve prendere due parametri da linea di comando
(in quest'ordine):
- numero N di linee da estrarre
- nome del file da elaborare

e stampare su standard output le ultime N righe del file. Se il
file ha meno di N righe, il programma stamperà tutte quelle che
ci sono.

Nota bene: NON va implementato invocando da Go il comando 'tail'
di sistema.

## Esempio esecuzione

(presuppone il 'tail.go' già compilato, 'uno.input' è presente
in questa directory)

$ ./tail 3 uno.input
remaining essentially unchanged. It was popularised in the 1960s
with the release of Letraset sheets containing Lorem Ipsum
passages, and more recently with desktop publishing software
like Aldus PageMaker including versions of Lorem Ipsum.
*/

package main

import(
  "fmt"
  "os"
  "bufio"
  "strconv"
  "bytes"
  "io"
)


func contaLinee(r io.Reader) (int, error){
  buffer := make([]byte, 32*1024)
  conta := 0
  speratore := []byte{'\n'}

  for {
    c, err := r.Read(buffer)
    conta += bytes.Count(buffer[:c], speratore)

    switch {
      case err == io.EOF:
        return conta, nil

      case err != nil:
        return conta, err
    }
  }
}

func main(){
  var righe int
  var r io.Reader

  if len(os.Args)!=3{
    fmt.Println("\tChiamare il programma e mettere un numero e un file, non più o meno di questo")
    os.Exit(1)
  }

  righe,err:=strconv.Atoi(os.Args[1])
  if err!=nil{
    fmt.Println(err)
    os.Exit(2)
  }

  file,err:=os.Open(os.Args[2])
  if err!=nil{
    fmt.Println(err)
    os.Exit(3)
  }
  r=file

  i,err:=contaLinee(r)
  if err!=nil{
    fmt.Println(err)
    os.Exit(3)
  }
  file.Close()

  file,err=os.Open(os.Args[2])
  if err!=nil{
    fmt.Println(err)
    os.Exit(3)
  }

  j:=0
  rigaDaPartire:=i-righe
  scanner:=bufio.NewScanner(file)
  if righe>=i{
    for scanner.Scan(){
      fmt.Println(scanner.Text())
    }
  }else{
    for scanner.Scan(){
      if j>=rigaDaPartire{
        fmt.Println(scanner.Text())
      }
      j++
    }
  }
  file.Close()
}
