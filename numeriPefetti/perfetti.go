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
  if len(os.Args)!=2{
    fmt.Println("mettere due argomenti")
    os.Exit(1)
  }

  num,_:=strconv.Atoi(os.Args[1])

  if sommaDivisori(num)==num{
    fmt.Printf("%d è perfetto",num)
  }else{
    fmt.Printf("%d non è pefetto",num)
  }
}
