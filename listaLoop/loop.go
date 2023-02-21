package main

import(
  "fmt"
)

type Nodo struct{
  valore int
  prec,succ *Nodo
}

type Lista struct{
  testa,coda *Nodo
}

func (l *Lista)inserisciValore(n int) bool{
  p:=&Nodo{valore: n,prec:nil,succ:nil}
  if l.testa==nil{
    l.testa=p

    return true
  }else if l.coda==nil{
    l.coda=p

    l.coda.succ=l.testa
    l.coda.prec=l.testa

    l.testa.succ=l.coda
    l.testa.prec=l.coda

    return true
  }else{
    t:=l.testa
    for t.succ!=l.testa{
      t=t.succ
    }
    t.succ=p
    p.succ=l.testa
    p.prec=t
    l.testa.prec=p
  }

  return true
}

func (l *Lista)stampaElenco(){
  p:=l.testa

  for p.succ!=l.testa{
    fmt.Printf("%d -->",p.valore)
    p=p.succ
  }
  fmt.Println("testa da capo")
}

func main(){
  var elenco Lista
  var n int

  for{
    fmt.Print("Valore (negativo per uscire): ")
    fmt.Scan(&n)
    if n<0{
      break
    }else{
      elenco.inserisciValore(n)
    }
  }

  elenco.stampaElenco()

}
