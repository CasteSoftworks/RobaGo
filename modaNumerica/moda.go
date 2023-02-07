package main

import(
  "fmt"
)

func moda(s []int) (m []int){
  check:=true
	mx:=0
	cont:=0

  for i,x:=range s{
    cont=0
    for j,y:=range s{
      if x==y&&i!=j{
        cont++
      }
    }
    if cont==mx{
      for _,e:=range m{
        if x==e{
          check=false
        }
      }
      if check{
        mx=cont
        m=append(m,x)
      }
    }else if cont>mx{
      for _,e:=range m{
        if x==e{
          check=false
        }
      }
      if check{
        mx=cont
        m=nil
        m=append(m,x)
      }
    }
  }

  return m
}

func main(){
  s:=[]int{5, -2, 3, -2, 3, -2, 3}
  fmt.Println(moda(s))
}
