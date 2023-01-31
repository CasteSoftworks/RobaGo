package main

import(
  "fmt"
  "os"
  "strings"
  "strconv"
)

type Data struct{
  giorno,mese,anno int
}

func ggDaEpochA(d Data) (gg int){
  gg+=d.giorno
  for i:=1;i<d.mese;i++{
    if i==2{
      if bisestile(d.anno){
        gg+=29
      }else{
        gg+=28
      }
    }else if i==4||i==6||i==9||i==11{
      gg+=30
    }else{
      gg+=31
    }
  }

  for i:=1970;i<d.anno;i++{
    if bisestile(i){
      gg+=366
    }else{
      gg+=365
    }
  }

  return gg
}

func bisestile(a int) bool{
  return ((a%4==0&&a%400!=0)||a%400==0)
}

func main(){
  if len(os.Args)!=3{
    fmt.Println("mettere due argomenti")
    os.Exit(1)
  }

  sliceDataN:=strings.Split(os.Args[1],"/")
  var dataN Data

  dataN.giorno,_=strconv.Atoi(sliceDataN[0])
  dataN.mese,_=strconv.Atoi(sliceDataN[1])
  dataN.anno,_=strconv.Atoi(sliceDataN[2])

  epochDaNascita:=ggDaEpochA(dataN)

  sliceDataO:=strings.Split(os.Args[2],"/")
  var dataO Data

  dataO.giorno,_=strconv.Atoi(sliceDataO[0])
  dataO.mese,_=strconv.Atoi(sliceDataO[1])
  dataO.anno,_=strconv.Atoi(sliceDataO[2])

  epochDaOra:=ggDaEpochA(dataO)

  fmt.Printf("sono passati %d giorni dalla tua nascita",(epochDaOra-epochDaNascita))

}
