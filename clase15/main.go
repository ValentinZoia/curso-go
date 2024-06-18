package main

import "fmt"

func mostrarAlgo(result string, c chan string) {
	c <- result
}

func main() {
	c := make(chan string)

	go mostrarAlgo("Hola channel", c)

	mjs := <-c
	fmt.Println(mjs)
}
