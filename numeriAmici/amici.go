package main

import(
  "fmt"
  "os"
  "strconv"
)

func sommaDivisori(n int) (somma int){
  for i:=1;i<n;i++{
    if n%i==0{
      somma+=i
    }else{
      continue
    }
  }

  return
}

func main(){
  if len(os.Args)!=3{
    fmt.Println("mettere tre argomenti")
    os.Exit(1)
  }

  num1,_:=strconv.Atoi(os.Args[1])
  num2,_:=strconv.Atoi(os.Args[2])

  if sommaDivisori(num1)==num2 && sommaDivisori(num2)==num1{
    fmt.Printf("%d e %d sono amici",num1,num2)
  }else{
    fmt.Printf("%d e %d non sono amici",num1,num2)
  }
}
