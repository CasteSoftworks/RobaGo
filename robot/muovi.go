package main

import(
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
)

type Comando struct{
  dir string //NORD SUD EST OVEST
  pas int //>0
}

func leggiComandi() (cmds []Comando){
  scan:=bufio.NewScanner(os.Stdin)
  for scan.Scan(){
    linea:=scan.Text()
    sliceC:=strings.Split(linea," ")

    n,err:=strconv.Atoi(sliceC[1])
    if err!=nil{
      break
    }

    if sliceC[0]!="NORD"&&sliceC[0]!="SUD"&&sliceC[0]!="EST"&&sliceC[0]!="OVEST"{
      break
    }
    var cmd Comando
    cmd.dir=sliceC[0]
    cmd.pas=n

    cmds=append(cmds,cmd)
  }

  return
}

func analizzaComandi(comandi []Comando) (mC map[string]int){
  mC=make(map[string]int)

  for i:=0;i<len(comandi);i++{
    mC[comandi[i].dir]+=comandi[i].pas
  }

  return
}

func inverso(cmd Comando) (inverso string){
  if cmd.dir=="NORD"{
    inverso+="SUD"
  }
  if cmd.dir=="SUD"{
    inverso+="NORD"
  }
  if cmd.dir=="EST"{
    inverso+="OVEST"
  }
  if cmd.dir=="OVEST"{
    inverso+="EST"
  }

  inverso+=" "

  inverso+=strconv.Itoa(cmd.pas)

  return
}

func main(){
  comandi:=leggiComandi()

  fmt.Println("Movimenti totali:")
  mCom:=analizzaComandi(comandi)
  for k,v:=range mCom{
    fmt.Println(k,v)
  }

  dirMax:=""
  pasMax:=0
  first:=true
  for k,v:=range mCom{
    if first{
      dirMax=k
      pasMax=v
      first=false
      continue
    }
    if v>pasMax{
      dirMax=k
      pasMax=v
    }
  }

  fmt.Printf("Direzione in cui il robot deve compiere il maggior numero totale di passi:\n%s\n",dirMax)

  fmt.Println("Comandi inversi:")
  for i:=len(comandi)-1;i>0;i--{
    fmt.Print(inverso(comandi[i]),", ")
  }
  fmt.Println(inverso(comandi[0]))
}
