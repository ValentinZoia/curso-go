package main

import "fmt"

type pc struct {
	ram        int
	disk       int
	procesador string
}

//ping()
func (myPc pc) ping() {
	fmt.Println(myPc.procesador, "Pong")
}

//implementacion de punteros, funcion que duplica el valor de la ram
func (myPc *pc) duplicateRAM() {
	myPc.ram = myPc.ram * 2

}

//modificar valor utilizando punteros
func cambiarMarcaProc(pc2 *pc, nuevaMarca string) {
	pc2.procesador = nuevaMarca
}

func main() {
	//Structs y Punteros

	a := 50
	b := &a

	fmt.Println(a, b, *b) //50 0xc00000a0a8 50

	*b = 100
	fmt.Println(*b) //100
	fmt.Println(a)  //100

	myPC := pc{
		ram:        16,
		disk:       200,
		procesador: "Ryzen",
	}

	fmt.Println(myPC) //{16 200 Ryzen}

	myPC.ping() //Ryzen Pong

	fmt.Println(myPC) //{16 200 Ryzen}
	myPC.duplicateRAM()

	fmt.Println(myPC) //{32 200 Ryzen}
	myPC.duplicateRAM()

	fmt.Println(myPC) //{64 200 Ryzen}

	fmt.Println("Procesador anterior", myPC.procesador) //Procesador anterior Ryzen
	cambiarMarcaProc(&myPC, "Intel")
	fmt.Println("Procesador nuevo", myPC.procesador) //Procesador nuevo Intel

}
