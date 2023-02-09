package main

import (
  "fmt"
  "math/rand"
  "time"
)
const baseNum=48

type Mano struct{
  c1,c2,c3 Carta
  puntTot int
}

type Carta struct{
  simbolo rune
  punteggio int
}

func generaCarta() (c Carta){
  rand.Seed(time.Now().UnixNano())

  pt:=rand.Intn(10)+1
  switch pt{
    case 1:
      c.simbolo='A'
      c.punteggio=11
    case 3:
      c.simbolo='3'
      c.punteggio=10
    case 10:
      c.simbolo='K'
      c.punteggio=4
    case 9:
      c.simbolo='Q'
      c.punteggio=3
    case 8:
      c.simbolo='J'
      c.punteggio=2
    default:
      c.simbolo=rune(baseNum+pt)
      c.punteggio=0
  }

  return
}

func assegnaMano() (m Mano){
  m.c1=generaCarta()
  time.Sleep(125*time.Millisecond)
  m.c2=generaCarta()
  time.Sleep(125*time.Millisecond)
  m.c3=generaCarta()
  time.Sleep(125*time.Millisecond)

  m.puntTot+=m.c1.punteggio
  m.puntTot+=m.c2.punteggio
  m.puntTot+=m.c3.punteggio

  return
}

func stampaMano(m Mano) string{
  return fmt.Sprintf("Mano dal valore di %d\n* carta: %c valore : %d\n* carta: %c valore : %d\n* carta: %c valore : %d",m.puntTot,m.c1.simbolo,m.c1.punteggio,m.c2.simbolo,m.c2.punteggio,m.c3.simbolo,m.c3.punteggio)
}

func main(){
  fmt.Println(stampaMano(assegnaMano()))
}
