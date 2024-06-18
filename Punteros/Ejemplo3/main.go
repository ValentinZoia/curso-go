package main

import "fmt"

// Definicion de la estructura Persona
type Persona struct {
	nombre string
	edad   int
}

// Funcion que modifica el nombre de una Persona
func cambiarNombre(p *Persona, nuevoNombre string) {
	p.nombre = nuevoNombre
}

func main() {
	persona := Persona{nombre: "Juan", edad: 30}
	fmt.Println("Antes de cambiar el nombre:", persona)

	cambiarNombre(&persona, "Carlos") // Pasamos la direccion de memoria de persona
	fmt.Println("Después de cambiar el nombre:", persona)
}
