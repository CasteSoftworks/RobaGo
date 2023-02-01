package main

import(
  "fmt"
)

func piuLungo(m [][]rune) (l int){
  cont:=0

  for i:=0;i<len(m);i++{

    for j:=0;j<len(m[i]);j++{
      if m[i][j]=='*'{
        cont++
      }
    }

    if l==0{
      l=cont
    }
    if l<cont{
      l=cont
    }

    cont=0
  }

  return

}

func main(){
  mat:=[][]rune{
    {' ','*','*',' '},
    {' ','*',' ',' '},
    {'*','*','*',' '},
    {'*','*','*','*'}}

  fmt.Println(piuLungo(mat))
}
