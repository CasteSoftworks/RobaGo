package main
//47
import(
  "fmt"
  "os"
  "strconv"
)

func primo(n int) bool{
  for i:=2;i<n;i++{
    if n%i==0{
      return false
    }
  }
  return true
}

func main(){
  if len(os.Args)<2{
    fmt.Println("almeno un argomento")
    os.Exit(1)
  }

  var mappaNumero map[int]bool

  mappaNumero=make(map[int]bool)

  for i:=1;i<len(os.Args);i++{
    n,_:=strconv.Atoi(os.Args[i])
    mappaNumero[n]=primo(n)
  }


  for key :=range mappaNumero{
    if mappaNumero[key]{
      fmt.Println(key, "è primo")
    }else{
      fmt.Println(key, "non è primo")

    }
  }

}
