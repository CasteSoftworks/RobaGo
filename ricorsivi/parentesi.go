package main

import(
  "fmt"
)

func contaPa(s string) int{
  p1,p2:=0,0
  ok1,ok2:= false, false

  for i:=0;i<len(s);i++{
    if s[i]=='('{
      p1=i
      ok1=true
      break
    }
  }

  for i:=len(s)-1;i>0;i--{
    if s[i]==')'{
      p2=i
      ok2=true
      break
    }
  }

  if ok1&&ok2{
    return (1+contaPa(s[p1+1:p2]))
  }else{
    return 0
  }
}

func main(){
  var s string

  s="(       (     (   (   (  ) ) )         )   )"

  fmt.Println(contaPa(s))
}
