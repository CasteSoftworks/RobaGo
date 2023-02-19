package main

import(
  "fmt"
  "os"
  "strconv"
  "bufio"
  "strings"
)

type Data struct{
  gg,mm,aa int
}

type Nodo struct{
  nome,cognome string
  matricola int
  nascita Data
  next *Nodo
}

type Lista struct{
  testa *Nodo
}

func (l *Lista)Inserisci(info []string) bool{
  matr,err:=strconv.Atoi(info[2])
  if err!=nil{
    fmt.Println("\t\t\terrore conversione")
    return false
  }

  sliceD:=strings.Split(info[3],"/")
  if len(sliceD)!=3{
    fmt.Println("\t\t\terrore nella data")
    return false
  }
  gg,err:=strconv.Atoi(sliceD[0])
  if err!=nil{
    fmt.Println("\t\t\terrore conversione")
    return false
  }
  mm,err:=strconv.Atoi(sliceD[1])
  if err!=nil{
    fmt.Println("\t\t\terrore conversione")
    return false
  }
  aa,err:=strconv.Atoi(sliceD[2])
  if err!=nil{
    fmt.Println("\t\t\terrore conversione")
    return false
  }
  dataNascita:=Data{gg:gg,mm:mm,aa:aa}

  lista:=&Nodo{nome: info[0],cognome: info[1],matricola: matr,nascita:dataNascita}

  if l.testa==nil{
    l.testa=lista
  }else{
    p:=l.testa
    for p.next!=nil{
      p=p.next
    }
    p.next=lista
  }

  return true
}

func main(){
  if len(os.Args)!=3{
    fmt.Println("mancano i nomi del file di ingresso e quello del file di uscita")
    os.Exit(1)
  }

  fIn,err:=os.Open(os.Args[1])
  if err!=nil{
    fmt.Println("errore in apertura file di ingresso")
    os.Exit(2)
  }

  studenti:=Lista{}

  scanner:=bufio.NewScanner(fIn)
  for scanner.Scan(){
    riga:=scanner.Text()
    sliceR:=strings.Split(riga,",")

    if len(sliceR)!=4{
      fmt.Println("\t\ttroppe/troppe poche informazioni sullo studente")
      continue
    }

    ok:=studenti.Inserisci(sliceR)
    if !ok{
      fmt.Println("\t\terrore in compilazione studente")
      continue
    }
  }

  fIn.Close()

  fOu,err:=os.OpenFile(os.Args[2],os.O_APPEND|os.O_WRONLY,0)
  if err!=nil{
    fmt.Println("errore in apertura file di uscita")
    os.Exit(2)
  }

  p:=studenti.testa
  for p!=nil{

    s:=fmt.Sprintf("Studente %s %s, matricola %d, nato il %d/%d/%d\n",p.nome,p.cognome,p.matricola,p.nascita.gg,p.nascita.mm,p.nascita.aa)

    n,erro:=fOu.WriteString(s)

    if erro!=nil{
      fmt.Println("errore in scrittura studente")
    }else{
      fmt.Printf("\t\tscritti %d bytes\n",n)
    }
    p=p.next
  }

  fOu.Close()
}
