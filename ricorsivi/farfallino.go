package main

import(
  "fmt"
)

func vocale(r rune) bool{
  if r=='a'||r=='e'||r=='i'||r=='o'||r=='u'||r=='A'||r=='E'||r=='I'||r=='O'||r=='U'{
    return true
  }

  return false
}

/*func f(s string) string{
  var t string

  for _,l:=range s{
    if vocale(l){
      t+=string(l)
      t+="f"
      t+=string(l)
    }else{
      t+=string(l)
    }
  }

  return t
}*/
func f(s string, j int){
  var t string

  if len(s)==0||j==len(s){
    return
  }else{
    t=string(s[j])
    if vocale(rune(s[j])){
      t+=("f"+string(s[j]))
    }
  }
  j++
  fmt.Printf("%s",t)
  f(s,j)
}

func main(){
  str:="supercalifragilistichespiralidoso"
  fmt.Print(str,"-->")
  f(str,0)
}
