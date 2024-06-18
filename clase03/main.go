package main

import "fmt"

func main() {
	//Operadores aritmeticos
	x := 10
	y := 50

	//Suma
	resultSuma := x + y
	fmt.Println(resultSuma) //60

	//Resta
	resultResta := x - y
	fmt.Println(resultResta) //-40

	//Multiplicacion
	resultMulti := x * y
	fmt.Println(resultMulti) //500

	//Division
	resultDivision := (y / x)
	fmt.Println(resultDivision) //5

	//Incrementar
	x++
	fmt.Println(x)

	//Decrementar
	x--
	fmt.Println(x)
}
