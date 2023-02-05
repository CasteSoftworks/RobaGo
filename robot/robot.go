/*
Creare un programma robot.go nel quale viene definito il tipo R obot
contenente le coordinate x e y del robot (interi) e la d irezione (un
tipo a piacere) che assumerà un valore corrispondente alla direzione
del robot (su/giù/destra/sinistra)
*/

package main

import(
  "fmt"
)

type Robot struct{
  x,y int
  dir int //0 su, 1 giù, 2 destra, 3 sinistra
}

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

func main(){
  droide,k:=costruisci(4,5,"su")
  
  if k{
    stato(droide)
    avanza(&droide)
    stato(droide)
    ruota(&droide,true)
    stato(droide)
    ruota(&droide,false)
    stato(droide)
  }
}
