package main

import(
  "fmt"
  "os"
  "strconv"
)

func scriviX(n int){
  h:=n-2
  k:=0

  for j:=0;j<n/2-1;j++{
    for i:=0;i<k;i++{
      fmt.Print(" ")
    }
    k++

    fmt.Print("*")

    for i:=0;i<h-2*k+1;i++{
      fmt.Print(" ")
    }

    fmt.Println("*")
  }

  for f:=0;f<k;f++{
    fmt.Print(" ")
  }
  if n%2!=0{
    fmt.Print("*")
  }
  fmt.Println("*")

  for j:=0;j<n/2-1;j++{
    for i:=k-1;i>0;i--{
      fmt.Print(" ")
    }
    k--

    fmt.Print("*")

    for i:=h-2*k-1;i>0;i--{
      fmt.Print(" ")
    }

    fmt.Println("*")
  }

}


func main(){
  n,_:=strconv.Atoi(os.Args[1])

  scriviX(n)
}
