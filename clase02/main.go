package main

import "fmt"

func main() {
	//Declarar Constantes
	const pi float64 = 3.14
	const pi2 = 3.1415
	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	//declarar variables enteras
	base := 12
	var altura int = 14
	var area int
	fmt.Println(base, altura, area)

	//Zero values
	var a int     //variables enteras
	var b float64 //variables con decimales
	var c string  //variables con textos
	var d bool    // variables true o false
	fmt.Println(a, b, c, d)

	//Ejercicio Area cuadrado
	const baseCuadrado int = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area =", areaCuadrado)
}
