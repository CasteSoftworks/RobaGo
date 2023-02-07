package main

import(
  "fmt"
)

type Data struct{
  gg,mm,aaaa uint
}
func bisestile(a uint) bool{
  return (a%4==0&&a%400!=0)||a%400==0
}
func creaData(gg,mm,aaaa uint) (d Data, ok bool){
  if aaaa<0{
    return
  }
  if mm<1||mm>12{
    return
  }
  if gg<1||gg>31{
    return
  }

  if mm==2{
    if bisestile(aaaa){
      if gg>29{
        return
      }
    }else{
      if gg>28{
        return
      }
    }
  }else if mm==4||mm==6||mm==9||mm==11{
    if gg>30{
      return
    }
  }

  d.gg=gg
  d.mm=mm
  d.aaaa=aaaa
  ok=true

  return
}

func aggiorna(d *Data){
  if d.mm==2{
    if bisestile(d.aaaa){
      if d.gg==29{
        d.gg=1
        d.mm+=1
      }else{
        d.gg+=1
      }
    }else{
      if d.gg==28{
        d.gg=1
        d.mm+=1
      }else{
        d.gg+=1
      }
    }
  }else if d.mm==4||d.mm==6||d.mm==9||d.mm==11{
    if d.gg==30{
      d.gg=1
      d.mm+=1
    }else{
      d.gg+=1
    }
  }else{
    if d.gg==31{
      if d.mm==12{
        d.gg=1
        d.mm=1
        d.aaaa+=1
      }else{
        d.gg=1
        d.mm+=1
      }
    }else{
      d.gg+=1
    }
  }
}

func main(){
  g,m,a:=29,2,2021
  d,k:=creaData(uint(g),uint(m),uint(a))

  if k{
    fmt.Println("creata:",d)
    aggiorna(&d)
    fmt.Println("data aggiornata:", d)
  }
}
