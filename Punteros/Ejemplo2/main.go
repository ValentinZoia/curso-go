/*
Ejemplo 2: Intercambiar dos valores usando punteros
*/

package main

import "fmt"

// Funcion que intercambia los valores de dos enteros usando punteros
func intercambiar(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	x := 1
	y := 2
	fmt.Println("Antes del intercambio:", x, y) //-->Antes del intercambio: 1 2

	intercambiar(&x, &y)                          // Pasamos las direcciones de memoria de x e y
	fmt.Println("Después del intercambio:", x, y) //-->Después del intercambio: 2 1
}
