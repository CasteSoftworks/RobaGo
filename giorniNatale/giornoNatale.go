//1/1/1970 era giovedì
package main

import(
  "fmt"
  "os"
  "strconv"
)

type Data struct{
  g,m,a int
}

func componi(sA []int)(dataGiorno map[Data]string){
  dataGiorno=make(map[Data]string)

  for _,a:=range sA{
    dN:=dataNatale(a)
    gio:=giorno(a)

    dataGiorno[dN]=gio
  }

  return
}

func dataNatale(a int) (natale Data){
  natale.g=25
  natale.m=12
  natale.a=a

  return
}

func giorno(a int) (gg string){
  g:=ggDaEpochNatale(a)

  switch g%7 {
  case 0:
    return "giovedì"
  case 1:
    return "venerdì"
  case 2:
    return "sabato"
  case 3:
    return "domenica"
  case 4:
    return "lunedì"
  case 5:
    return "martedì"
  default:
    return "mercoledì"
  }
}

func ggDaEpochNatale(a int) (g int){
  for i:=1970;i<a;i++{
    g+=lungAnno(i)
  }

  if bisestile(a){
    g+=359
  }else{
    g+=358
  }

  return
}

func lungAnno(a int) int{
  if bisestile(a){
    return 366
  }
  return 365
}

func bisestile(a int) bool{
  return ((a%4==0&&a%400!=0)||a%400==0)
}

func main(){
  if len(os.Args)<2{
    fmt.Println("almeno un argomento")
    os.Exit(1)
  }

  var sliceNum []int
  sliceNum=make([]int,1)

  for _,a:=range os.Args[1:]{
    anno,_:=strconv.Atoi(a)
    sliceNum=append(sliceNum,anno)
  }

  sliceNum=sliceNum[1:]

  fmt.Println(componi(sliceNum))

}
