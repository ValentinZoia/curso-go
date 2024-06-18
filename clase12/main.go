package main

import (
	"fmt"
	"math"
)

// Definición de la interfaz Forma
type Forma interface {
	Area() float64
}

// Definición de la estructura Círculo
type Circulo struct {
	Radio float64
}

// Implementación del método Area para Círculo
func (c Circulo) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}

// Definición de la estructura Rectángulo
type Rectangulo struct {
	Ancho, Alto float64
}

// Implementación del método Area para Rectángulo
func (r Rectangulo) Area() float64 {
	return r.Ancho * r.Alto
}

func main() {
	// Crear instancias de Círculo y Rectángulo
	circulo := Circulo{Radio: 5}
	rectangulo := Rectangulo{Ancho: 4, Alto: 6}

	// Crear una lista de Forma
	formas := []Forma{circulo, rectangulo}

	// Iterar sobre la lista y calcular el área de cada forma
	for _, forma := range formas {
		fmt.Printf("Área: %.2f\n", forma.Area())
	}

	//lista de interfaces
	myInterface := []interface{}{"Hola", 12, 4.90}
	fmt.Println(myInterface...)
}
