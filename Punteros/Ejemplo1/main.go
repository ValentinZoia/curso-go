/*
Ejemplo 1: Modificar el valor de una variable usando punteros
*/

package main

import "fmt"

// Funcion que recibe un puntero a un entero y modifica su valor
func incrementarValor(val *int) {
	*val += 10 // Incrementa el valor al que apunta el puntero por 10
}

func main() {
	x := 5
	fmt.Println("Antes de incrementar:", x) //-->Antes de incrementar: 5

	incrementarValor(&x)                      // Pasamos la direccion de memoria de x
	fmt.Println("Después de incrementar:", x) //-->Después de incrementar: 15
}
