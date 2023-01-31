/*
Programma che, dopo avere inserito da linea di comando la velocità di diaframma
e l'intensità luminosa comunica all'utente l'apertura necessaria affinché la
ripresa venga esposta correttamente
*/

package main

import(
  "fmt"
  "os"
  "math"
  "strconv"
)

func luxToEV(l int) (ev int){
  return int(math.Log2((float64(l))/2))
}

func calculateAperture(iso int, shutSpeed int, ev int) (fStop float64){
  shutSpeedTrue:=1/float64(shutSpeed)
  return (math.Sqrt((shutSpeedTrue*(float64(iso)*math.Pow(2,float64(ev))))/100))
}

func main(){
  var shutterSpeed, lux, ev, iso int
  if len(os.Args)<4{
    fmt.Fprintf(os.Stderr,"METTI UNA VELOCITA' di DIAFRAMMA, UN VALORE IN LUX E UN VALORE ISO\n")
    os.Exit(1)
  }

  shutterSpeed,err:=strconv.Atoi(os.Args[1])
  if err!=nil{
    fmt.Fprintf(os.Stderr,"\t\terrore di conversione\n")
    os.Exit(2)
  }
  lux,err=strconv.Atoi(os.Args[2])
  if err!=nil{
    fmt.Fprintf(os.Stderr,"\t\terrore di conversione\n")
    os.Exit(2)
  }
  iso,err=strconv.Atoi(os.Args[3])
  if err!=nil{
    fmt.Fprintf(os.Stderr,"\t\terrore di conversione\n")
    os.Exit(2)
  }

  ev=luxToEV(lux)
  fmt.Printf("EV:%23d\nVELOCITA' DIAFRAMMA:%6d\nISO:%22d",ev,shutterSpeed,iso)

  fmt.Printf("\n\n\t\tAPERTURA CONSIGLIATA:\t%f",calculateAperture(iso,shutterSpeed,ev))

}
