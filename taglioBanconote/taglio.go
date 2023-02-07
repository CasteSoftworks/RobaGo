package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Errore")
	}

	n,_:=strconv.Atoi(os.Args[1])
	nf:=float64(n)

	fmt.Println("Taglio: 100 Quantità:", int(nf)/100)
	nf=nf-(nf-float64((n%100)))
	n=int(nf)

	fmt.Println("Taglio: 50 Quantità:", int(nf)/50)
	nf=nf-(nf - float64((n%50)))
	n=int(nf)

	fmt.Println("Taglio: 20 Quantità:", int(nf)/20)
	nf=nf-(nf-float64((n%20)))
	n=int(nf)

	fmt.Println("Taglio: 10 Quantità:", int(nf)/10)
	nf=nf-(nf-float64((n%10)))
	n=int(nf)

	fmt.Println("Taglio: 5 Quantità:", int(nf)/5)
	nf=nf-(nf-float64((n%5)))
	n=int(nf)

	fmt.Println("Taglio: 2 Quantità:", int(nf)/2)
	nf=nf-(nf-float64((n%2)))
	n=int(nf)

	fmt.Println("Taglio: 1 Quantità:", int(nf)/1)
	nf=nf-(nf-float64((n%1)))
	n=int(nf)
}
