package main

import "fmt"

// Calculacion Numero E

func main(){
	var ite int 
	var calcPrev float64 = 1
	var valorAct float64 = 1
	fmt.Scan(&ite)
	for i := 1; i <= ite; i++ {
		calcPrev = (calcPrev * float64(i))
		valorAct = valorAct + (1 / calcPrev)
	}

	fmt.Println(valorAct)
}