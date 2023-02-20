package main

import(
  "fmt"
  "os"
  "strconv"
)

func tartaglia(k,n int) int{
  if k==0||k==n{
    return 1
  }else if n<k||n<0||k<0{
    return 0
  }else{
    return tartaglia(k-1,n-1)+tartaglia(k,n-1)
  }
}

func main(){
  iter,err:=strconv.Atoi(os.Args[1])
  if err!=nil{
    fmt.Println("errore")
    os.Exit(1)
  }

  for n:=0; n<=iter; n++{
    for k:=0; k<=n; k++{
      fmt.Printf("%3d ", tartaglia(k, n))
    }
    fmt.Println()
  }
}
