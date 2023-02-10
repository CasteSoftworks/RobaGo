/*
Scrivere un programma 'temperature.go' che legge da standard
input una sequenza (terminata da EOF) di almeno una temperatura
e stampa la temperatura minima e il numero di volte che è stata
registrata, la temperatura massima e il numero di volte che è
stata registrata, le temperature che si sono verificate con la
massima frequenza.

Esempio
-------
su input
2 3 4 -2 2 6 12 8 3 1 -3
0 -2 5 2 4 1 -1 0 7 11 8 3

produce in output

max : 12, 1 volte; min : -3, 1 volte
-3:1 -2:2 -1:1 0:2 1:2 2:3 3:3 4:2 5:1 6:1 7:1 8:2 11:1 12:1
temperature [2 3] con frequenza 3, la massima frequenza
*/

package main

import(
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
)

func trovaEstremi(s []float64) (float64, float64, int, int){
  var max,min float64
  var nMax,nMin int

  primo:=true
  for _,e:=range s{
    if primo{
      max=e
      nMax=1
      primo=false
    }else{
      if max<e{
        max=e
        nMax=1
      }else if max==e{
        nMax++
      }else{
        continue
      }
    }
  }

  primo=true
  for _,e:=range s{
    if primo{
      min=e
      nMin=1
      primo=false
    }else{
      if min>e{
        min=e
        nMin=1
      }else if min==e{
        nMin++
      }else{
        continue
      }
    }
  }

  return max,min,nMax,nMin
}

func frequenza(s []float64) (mF map[float64]int){
  mF=make(map[float64]int)

  for i:=0;i<len(s);i++{
    mF[s[i]]=0
    for j:=0;j<len(s);j++{
      if s[i]==s[j]{
        mF[s[i]]++
      }
    }
  }

  return
}

func soloDecimaliNecessari(f float64) string{
  return strconv.FormatFloat(f, 'f', -1, 64)
}

func massimaFrequenza(m map[float64]int) (sM []string,fM int){
  primo:=true
  for k,v:=range m{
    if primo{
      fM=v
      sM=append(sM,soloDecimaliNecessari(k))
      primo=false
      continue
    }
    if fM<v{
      fM=v
      sM=nil
      sM=append(sM,soloDecimaliNecessari(k))
    }else if fM==v{
      sM=append(sM,soloDecimaliNecessari(k))
    }else{
      continue
    }
  }

  return
}

func main(){
  var sliceTemp []float64

  scanner:=bufio.NewScanner(os.Stdin)
  for scanner.Scan(){
    riga:=scanner.Text()
    sT:=strings.Split(riga," ")

    for _,e:=range sT{
      fT,_:=strconv.ParseFloat(e,64)
      sliceTemp=append(sliceTemp,fT)
    }
  }

  max,min,nMax,nMin:=trovaEstremi(sliceTemp)

  fmt.Printf("max : %s, %d volte; min : %s, %d volte\n",soloDecimaliNecessari(max),nMax,soloDecimaliNecessari(min),nMin)

  mappaTemp:=frequenza(sliceTemp)
  for k,v:=range mappaTemp{
    fmt.Printf("%s:%d ",soloDecimaliNecessari(k),v)
  }
  fmt.Println()

  sM,fM:=massimaFrequenza(mappaTemp)
  fmt.Printf("temperature %v con frequenza %d, la massima frequenza\n",sM,fM)
}
