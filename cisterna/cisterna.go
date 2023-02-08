package main

import (
	"fmt"
  "os"
  //"errors"
  "strconv"
)

type Cisterna struct{
  x,y,z float64 //x larghezza , y lunghezza, z altezza
  liv float64 //cm
}
/*
      |    /
      |z /
______|/y
x
*/

func variazione(cisterna *Cisterna, lt int) error{
  var cT Cisterna

  if lt>0{
    cT.x=cisterna.x
    cT.y=cisterna.y
    cT.z=cisterna.z
    cT.liv=cisterna.z

    lDisp:=contenuto(cT)-contenuto(*cisterna)
    if lt>lDisp{
      err:=fmt.Errorf("puoi mettere max %d litri",lDisp)

      return err
    }else{
      cisterna.liv+=((1000*float64(lt)/(cisterna.x*cisterna.y)))
    }
  }else if lt<0{
    lPrima:=contenuto(*cisterna)
    if lPrima+lt<0{
      err:=fmt.Errorf("puoi prendere max %d litri",lPrima)

      return err
    }else{
      cisterna.liv+=((1000*float64(lt)/(cisterna.x*cisterna.y)))
    }
  }

  return nil
}

func contenuto(cisterna Cisterna) (litri int){
  if cisterna.liv!=0{
    litri=int(((cisterna.x*cisterna.y)*cisterna.liv)/1000)
  }

  return
}

func (c *Cisterna)String() string{
  return fmt.Sprintf("cisterna %.2f cm x %.2f cm x %.2f cm\nlivello attuale: %.2f cm, litri %d",c.x,c.y,c.z,c.liv,contenuto(*c))
}

func main(){
  if len(os.Args)!=4{
    fmt.Println("ERRORE")
    os.Exit(1)
  }

  lar,err:=strconv.Atoi(os.Args[1])
  if err!=nil||lar<0{
    fmt.Println("tutti gli argomenti devono essere >0")
    os.Exit(2)
  }
  lun,err:=strconv.Atoi(os.Args[2])
  if err!=nil||lun<0{
    fmt.Println("tutti gli argomenti devono essere >0")
    os.Exit(2)
  }
  alt,err:=strconv.Atoi(os.Args[3])
  if err!=nil||alt<0{
    fmt.Println("tutti gli argomenti devono essere >0")
    os.Exit(2)
  }

  var c Cisterna
  c.x=float64(lar)
  c.y=float64(lun)
  c.z=float64(alt)
  c.liv=0

  aggRim:=0
  for{
    s:=c.String()
    fmt.Println(s)

    fmt.Scan(&aggRim)

    if aggRim==9999{
      break
    }
    
    errore:=variazione(&c,aggRim)
    if errore!=nil{
      fmt.Println(errore)
    }
  }
}
