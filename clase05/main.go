package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("chi")
	} else {
		fmt.Println("no")
	}

	//With and
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("es verdad")
	}

	//Whith or - ||
	if valor1 == 34 || valor2 == 2 {
		fmt.Println("Una condicion o ambas se cumplen")
	}

	//Covertir texto a numero
	//dos retornos value(valor convertido) y err(error)
	// si no hay error. err sera nil
	value, err := strconv.Atoi("54")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value:", value)
	//value, err := strconv.Atoi("safds") -->strconv.Atoi: parsing "safds": invalid syntax

	//Reto: Reakizar una funcion que valide si un numero es par o no.
	fmt.Println(esPar(220))
}

func esPar(x int) string {
	if x%2 != 0 {
		return "es impar"
	}
	return "es par"
}
