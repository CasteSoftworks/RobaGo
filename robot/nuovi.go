/*Estendere il programma robot.go*/

package main

import(
  "fmt"
)

type Robot struct{
  x,y int
  dir int //0 su, 1 giù, 2 destra, 3 sinistra
}

//frecce commentate perché non stampano su winzozz
const su="u"//"⬆"
const giu="d"//"⬇"
const dx="r"//"➡"
const sx="l"//"⬅"
const vuota="□"

/*Definire (inferendone il prototipo) la funzione c ostruisci che, dati
due interi ed una stringa, restituisca una variabile di tipo R obot ed
un valore booleano. Se tutti i dati passati sono corretti, il valore
booleano ritornato sarà t rue, e la variabile di tipo R obot restituita
avrà la cordinata x corrispondente al primo valore passato, la
coordinata y corrispondente al secondo valore passato, e la d irezione
stabilita in base alla stringa passata come terzo valore.
Specificatamente, se la stringa è "su", "giù", "destra" o "sinistra",
la direzione del robot avrà il valore corrispondente. Se la stringa è
diversa dai valori specificati sopra, il valore booleano restituito
sarà falso. In questo caso non ci interessa come sarà la variabile
robot. Es:
  r2d2,ok := costruisci(5,5,"su") //ok sarà true
  c3po,ko := costruisci(0,0,"rotto") //ko sarà false*/
func costruisci(x,y int, s string) (r Robot,ok bool){
  switch s {
  case "su","up":
    r.dir=0
  case "giù","giu","down":
    r.dir=1
  case "destra","dx","right":
    r.dir=2
  case "sinistra","sx","left":
    r.dir=3
  default:
    return
  }

  r.x=x
  r.y=y

  ok=true
  return
}

/*Definire (inferendone il prototipo) la funzione s tato che, dato un
parametro di tipo R obot stampi le sue coordinate e la sua direzione.
Si assuma che solo i Robot costruiti correttamente verranno passati.
Es:
  stato(r2d2)
restituirà:
  {5 5 su}*/
func stato(r Robot){
  fmt.Printf("{%d %d",r.x,r.y)
  switch r.dir {
  case 0:
    fmt.Printf(" su}\n")
  case 1:
    fmt.Printf(" giù}\n")
  case 2:
    fmt.Printf(" destra}\n")
  case 3:
    fmt.Printf(" sinistra}\n")
  }
}

/*Definire (inferendone il prototipo) la funzione a vanza che, dato un
puntatore a un R obot ( si assuma corretto), faccia avanzare il robot
nella direzione corrente di una posizione. Si assuma che la coordinata
x vada da sinistra verso destra e la coordinata y da su verso giù.
Quindi se la direzione del robot è su, avanzare decrementerà la
coordinata y di 1. Es:
  avanza(&r2d2)
  stato(r2d2)
restituirà:
  {5 4 su}*/
func avanza(r *Robot){
  switch r.dir {
  case 0:
    r.y--
  case 1:
    r.y++
  case 2:
    r.x++
  case 3:
    r.x--
  }
}

/*Definire (inferendone il prototipo) la funzione r uota che, dato un
puntatore a un R obot ( si assuma corretto) ed un valore booleano che
specifichi il senso della rotazione (se true orario, altrimenti
antiorario), giri il robot nel senso indicato. Es:
  ruota(&r2d2, True)
  stato(r2d2)
restituirà:
  {5 4 destra}*/
func ruota(r *Robot, senso bool){
  if senso{
    switch r.dir {
    case 0:
      r.dir=2
    case 1:
      r.dir=3
    case 2:
      r.dir=1
    case 3:
      r.dir=0
    }
  }else{
    switch r.dir {
    case 0:
      r.dir=3
    case 1:
      r.dir=2
    case 2:
      r.dir=0
    case 3:
      r.dir=1
    }
  }
}

/*la funzione griglia che, data una slice di Robot (si assuma corretti e non sovrapposti),
e quattro valori interi (xmin, xmax, ymin, ymax), mostri a schermo una
griglia che si estende tra le coordinate x e y minime e massime
specificate (incluse). Per ciascuna casella nella quale non vi è un
robot presente verrà stampato il carattere "□" (valore intero 9633).
Per ciascuna casella nella quale vi è presente un robot verrà stampato
un carattere che dipenderà dalla direzione del robot: ⬆ (valore intero
11014) se su, ⬇ (valore intero 11015) se giù, ⬅ (valore intero 11013)
se sinistra, ➡ (valore intero 10145) se destra. Es:
  r1,_ := costruisci(1,1,"destra")
  r2,_ := costruisci(2,2,"giù")
  r3,_ := costruisci(2,5,"su")
  r4,_ := costruisci(2,7,"su")
  robots := []Robot{r1,r2,r3,r4}
  griglia(robots, 0, 4, 0, 3)
restituirà:
  □□□□□
  □➡□□□
  □□⬇□□
  □□□□□*/
func griglia(rs []Robot, xmin, xmax, ymin, ymax int){
  pres:=false
  for i:=ymin;i<ymax;i++{
    for j:=xmin;j<xmax;j++{
      for _,r :=range rs{
        if r.x==j && r.y==i{
          pres=true
          switch r.dir {
          case 0:
            fmt.Print(" ",su)
          case 1:
            fmt.Print(" ",giu)
          case 2:
            fmt.Print(" ",dx)
          case 3:
            fmt.Print(" ",sx)
          }
          break
        }
      }
      if pres{
        j++
        pres=false
      }
      fmt.Print(" ",vuota)
    }
    fmt.Println()
  }
}

/*funzione vista che, data una slice di robot (si assuma corretti e
non sovrapposti), restituisca una slice contenente tutte le coppie di
robot che si vedono tra di loro (le coppie sono slice di due elementi).
Due robot si vedono tra di loro se sono sulla stessa linea (hanno la
stessa coordinata x o y) e sono direzionati uno nel verso dell'altro.
ATTENZIONE: un robot vede tutti i robot nella stessa linea verso i
quali è direzionato e che sono direzionati verso di lui. Ad esempio,
considerando la slice robots dell'esercizio precedente:
  c := vista(robots)
  fmt.Println(c)
stamperà qualcosa di questo genere (Nota: non è richiesta la stampa
perfetta ma solo che la slice restituita sia corretta):
  [[{2 2 giù} {2 5 su}] [{2 2 giù} {2 7 su}]]*/
func vista(sR []Robot) (vede [][]Robot){
  for _,r:=range sR{
    for i:=0;i<len(sR);i++{
      if r!=sR[i]{
        if r.x==sR[i].x || r.y==sR[i].y{
          if r.x>sR[i].x{
            switch r.dir {
              case 3:
                if sR[i].dir==2{
                  vdn:=[]Robot{r,sR[i]}
                  vede=append(vede,vdn)
                }else{
                  continue
                }
              default:
                continue
            }
          }else if r.x<sR[i].x{
            switch r.dir {
              case 2:
                if sR[i].dir==3{
                  vdn:=[]Robot{r,sR[i]}
                  vede=append(vede,vdn)
                }else{
                  continue
                }
              default:
                continue
            }
          }else if r.y>sR[i].y{
            switch r.dir {
              case 0:
                if sR[i].dir==1{
                  vdn:=[]Robot{r,sR[i]}
                  vede=append(vede,vdn)
                }else{
                  continue
                }
              default:
                continue
            }
          }else if r.y<sR[i].y{
            switch r.dir {
              case 1:
                if sR[i].dir==0{
                  vdn:=[]Robot{r,sR[i]}
                  vede=append(vede,vdn)
                }else{
                  continue
                }
              default:
                continue
            }
          }
        }
      }else{
        continue
      }
    }
  }
  return
}

/*Estendere il programma r obot.go, scrivendo la funzione main in modo
tale che richieda all'utente di inserire dei nuovi robot da standard
input, uno alla volta. Il robot dovrà essere inserito specificando le
coordinate x e y come interi, e la direzione come stringa. I valori
possibili per la direzione sono su/giù/destra/sinistra. Qualora un
robot sia già presente alle coordinate specificate, il nuovo robot non
dovrà essere inserito, il sistema scriverà "occupato" e chiederà un
nuovo inserimento. Inserendo un valore di direzione non valido si
terminerà l'input. Infine, dovrà essere stampata la lista dei robot
inseriti, es (le righe indentate sono inserite dall'utente):
  Inserisci un robot:
    2 2 su
  Inserisci un robot:
    0 1 destra
  Inserisci un robot:
    0 1 destra
      occupato
  Inserisci un robot:
    0 0 fine
  Robot inseriti:
  {0 1 destra}
  {2 2 su}*/
func main(){
  var direzione string
  var x,y int
  var robots []Robot

  primo:=true
  occupato:=false
  for {
    fmt.Println("Inserisci un robot:")
    fmt.Scan(&x,&y,&direzione)
    if direzione=="fine"{
      break
    }
    if primo{
      rb,k:=costruisci(x,y,direzione)
      if !k{
        fmt.Println("\t\t\tERRORE")
        continue
      }
      robots=append(robots,rb)
      primo=false
    }else{
      for _,r:=range robots{
        if x==r.x&&y==r.y{
          fmt.Println("occupato")
          occupato=true
          break
        }
      }
      if occupato{
        occupato=false
        continue
      }
      rb,k:=costruisci(x,y,direzione)
      if !k{
        fmt.Println("\t\t\tERRORE")
        continue
      }
      robots=append(robots,rb)
    }
  }

  fmt.Println("Robot inseriti:")
  for _,r:=range robots{
    stato(r)
  }

  fmt.Println("Griglia:")
  griglia(robots,0,20,0,20)

  fmt.Println("Si vedono:")
  fmt.Println(vista(robots))

}
