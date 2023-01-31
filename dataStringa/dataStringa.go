package main

import(
  "fmt"
  "os"
)

func bisestile(a int) bool{
  return ((a%4==0&&a%400!=0)||a%400==0)
}

func main(){
  var g,m,a int
  fmt.Print("inserire una data (gg/mm/aaaa): ")
  fmt.Scan(&g,&m,&a)

  if g<0{
    fmt.Println("errore")
    os.Exit(1)
  }

  if m<0||m>12{
    fmt.Println("errore")
    os.Exit(1)
  }

  if bisestile(a){
    if m==2{
      if g>29{
        fmt.Println("errore")
        os.Exit(1)
      }
    }else if m==4||m==6||m==9||m==11{
      if g>30{
        fmt.Println("errore")
        os.Exit(1)
      }
    }else{
      if g>31{
        fmt.Println("errore")
        os.Exit(1)
      }
    }
  }else{
    if m==2{
      if g>28{
        fmt.Println("errore")
        os.Exit(1)
      }
    }else if m==4||m==6||m==9||m==11{
      if g>30{
        fmt.Println("errore")
        os.Exit(1)
      }
    }else{
      if g>31{
        fmt.Println("errore")
        os.Exit(1)
      }
    }
  }

  fmt.Printf("%s",fmt.Sprintf("%d/%d/%d",g,m,a))


}
