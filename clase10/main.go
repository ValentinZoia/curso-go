package main

import "fmt"

type Camiseta struct {
	colorPrimario   string
	colorSecundario string
	marca           string
	numero          int
}

// Implementación del método String para la estructura Persona
func (c Camiseta) String() string {
	return fmt.Sprintf("Camiseta color %s y %s, marca %s y numero %d ", c.colorPrimario, c.colorSecundario, c.marca, c.numero)
}

func main() {
	//Structs: La forma de hacer clases en Go

	//Stringers:personalizar el output de Structs

	miCasaca := Camiseta{
		colorPrimario:   "azul",
		colorSecundario: "rojo",
		marca:           "nike",
		numero:          11,
	}

	fmt.Println(miCasaca) //Camiseta color azul y rojo, marca nike y numero 11
}
