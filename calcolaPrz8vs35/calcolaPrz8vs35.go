package main

import (
    "fmt"
    "os"
    "strconv"
    "example.com/caste/calcola8mm"
    "example.com/caste/calcola35mm"
)

func main() {
    if len(os.Args)<6{
      fmt.Fprintf(os.Stderr,"troppi pochi parametri\n")
      os.Exit(1)
    }
    ore,err:=strconv.Atoi(os.Args[1])
    if err!=nil{
      fmt.Fprintf(os.Stderr,"errore di conversione\n")
      os.Exit(2)
    }
    min,err:=strconv.Atoi(os.Args[2])
    if err!=nil{
      fmt.Fprintf(os.Stderr,"errore di conversione\n")
      os.Exit(2)
    }
    sec,err:=strconv.Atoi(os.Args[3])
    if err!=nil{
      fmt.Fprintf(os.Stderr,"errore di conversione\n")
      os.Exit(2)
    }
    prezzo8,err:=strconv.ParseFloat(os.Args[4],64)
    if err!=nil{
      fmt.Fprintf(os.Stderr,"errore di conversione\n")
      os.Exit(2)
    }
    prezzo35,err:=strconv.ParseFloat(os.Args[5],64)
    if err!=nil{
      fmt.Fprintf(os.Stderr,"errore di conversione\n")
      os.Exit(2)
    }
    costo8,quanti8:=calcola8mm.Calcola(ore,min,sec,prezzo8)
    fmt.Printf("prezzo di %d:%d:%d di girato a %.2f € a rotolo 8mm: %.2f € per %d rotoli\n",ore,min,sec,prezzo8,costo8,quanti8)

    costo35,quanti35:=calcola35mm.Calcola(ore,min,sec,prezzo35)
    fmt.Printf("prezzo di %d:%d:%d di girato a %.2f € a rullino 35mm: %.2f € per %d rullini\n",ore,min,sec,prezzo35,costo35,quanti35)

    costo8+=float64(quanti8)*28
    costo35+=float64(quanti35)*8
    fmt.Printf("\n\nconsiderando 28€ per rullo 8mm e 8€ per ogni rullino 35mm\nCosto rotoli+sviluppo 8mm: %.2f\nCosto rullini+sviluppo 35mm: %.2f\n",costo8,costo35)

    costo8+=float64(quanti8)*10
    costo35+=float64(quanti35)*10
    fmt.Printf("\n\nconsiderando 10€ per rullo/rullino di scan\nCosto rotoli+sviluppo+scan 8mm: %.2f\nCosto rullini+sviluppo+scan 35mm: %.2f\n",costo8,costo35)


}
