/* date due slice di stringhe, una con numeri di telefoni sporchi (--> con caratteri non
numerici presenti) ed una con numeri di telefono puliti (--> con caratteri solo numerici)
restituisce quanti dei numeri presenti in dirty sono presenti anche in clean*/

package main

import(
  "fmt"
)
func pulisci(s string) (n string){
  for _,e:=range s{

    if e>='0' && e<='9'{
      n+=string(e)
    }
  }

  return
}

func f(s,t []string) (q int){

  for _,v:=range s{
    num1:=pulisci(v)

    for _,k:=range t{
      num2:=pulisci(k)

      if num1==num2{
        q++
      }
    }
  }

  return
}

func main(){
  s:=[]string{"38848c11920","36t68975f430","3234678904"}
  t:=[]string{"38848c11920","3884811920","3668975f530","3234678905","3668975f430","3234678905"}

  fmt.Println(f(s,t))
}
