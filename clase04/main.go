package main

import "fmt"

func main() {
	nombre := "Pedro"
	edad := 23
	fmt.Printf("%s tiene %d años \n", nombre, edad) //Pedro tiene 23 años
	message := fmt.Sprintf("%s tiene %d años", nombre, edad)
	fmt.Println(message)                                   //Pedro tiene 23 años
	fmt.Printf("la variable nombre es de tipo %T", nombre) //la variable nombre es de tipo string
}
