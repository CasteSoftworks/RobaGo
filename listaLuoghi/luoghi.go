package main

import(
  "fmt"
  "strconv"
  "strings"
  "time"
)

type Data struct{
  gg,mm,aa int
}

type Luogo struct{
  nome,citta,provincia,regione string
  visitato Data
  next *Luogo
}

type Lista struct{
  testa *Luogo
}

func disegnaMenu(){
  fmt.Print("----------------------\n")
  fmt.Print("| I: inserisci luogo |\n")
  fmt.Print("| C: cerca luogo     |\n")
  fmt.Print("| E: elimina luogo   |\n")
  fmt.Print("| S: stampa luoghi   |\n")
  fmt.Print("| 0: esci            |\n")
  fmt.Print("----------------------\n")
}

func (l *Lista)inserisciLuoghi() bool{
  var n,c,p,r,dta string

  fmt.Println("Inserisci nome, città, provincia e regione (separate da spazio, se il nome contiene spazi, usare _):")
  fmt.Scan(&n,&c,&p,&r)

  fmt.Println("Inserisci ultima visita effetuata (gg/mm/aaaa):")
  fmt.Scan(&dta)
  slD:=strings.Split(dta,"/")
  g,err:=strconv.Atoi(slD[0])
  if err!=nil{
    return false
  }
  m,err:=strconv.Atoi(slD[1])
  if err!=nil{
    return false
  }
  a,err:=strconv.Atoi(slD[2])
  if err!=nil{
    return false
  }
  vis:=Data{gg:g,mm:m,aa:a}

  posto:=&Luogo{nome: n,citta: c,provincia: p,regione: r, visitato: vis}

  if l.testa==nil{
    l.testa=posto
  }else{
    p:=l.testa
    for p.next!=nil{
      p=p.next
    }
    p.next=posto
  }

  return true
}

func (l *Lista)stampaLuoghi(){
  p:=l.testa

  for p!=nil{
    fmt.Printf("%s, %s (%s), %s visitato il %d/%d/%d\n",p.nome,p.citta,p.provincia,p.regione,p.visitato.gg,p.visitato.mm,p.visitato.aa)
    p=p.next
  }
}

func (l *Lista)cercaLuogo(nome string) (Luogo,bool){
  p:=l.testa

  for p!=nil{
    if p.nome==nome{
      return *p,true
    }
    p=p.next
  }

  return *p,false
}

func (l *Lista)eliminaLuogo(nome string){

  if l.testa.nome==nome{
    l.testa=l.testa.next
    return
  }
  var prec *Luogo=nil
  curr:=l.testa
  for curr!=nil &&curr.nome!=nome{
    prec=curr
    curr=curr.next
  }
  prec.next=curr.next
  return
}

func main(){
  var cmd string
  var elenco Lista

  for {
    disegnaMenu()
    fmt.Print("Comando > ")
    fmt.Scan(&cmd)
    if cmd=="0"{
      break
    }else if cmd=="I"{
      ok:=elenco.inserisciLuoghi()
      if !ok{
        fmt.Println("Errore in inserimento, salto a prossimo")
        continue
      }
    }else if cmd=="S"{
      elenco.stampaLuoghi()
    }else if cmd=="C"{
      var nome string
      fmt.Print("\tNome luogo da cercare: ")
      fmt.Scan(&nome)
      p,ok:=elenco.cercaLuogo(nome)
      if ok{
        fmt.Println("trovato")
        fmt.Printf("%s, %s (%s), %s visitato il %d/%d/%d\n",p.nome,p.citta,p.provincia,p.regione,p.visitato.gg,p.visitato.mm,p.visitato.aa)
      }else{
        fmt.Println("non trovato")
      }
    }else if cmd=="E"{
      var nome string
      fmt.Print("\tNome luogo da eliminare: ")
      fmt.Scan(&nome)
      _,ok:=elenco.cercaLuogo(nome)
      if ok{
        fmt.Println("luogo trovato, procedo alla eliminazione")
        time.Sleep(50 * time.Millisecond)
        fmt.Print(".")
        elenco.eliminaLuogo(nome)
        time.Sleep(50 * time.Millisecond)
        fmt.Print(".")
        time.Sleep(50 * time.Millisecond)
        fmt.Print(".")
        fmt.Println("Fatto")
      }else{
        fmt.Println("luogo non trovato")
      }
    }else{
      fmt.Println("comando non riconosciuto")
    }
  }
}
