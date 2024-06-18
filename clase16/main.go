//Sefuimos con cocncurrencia.
// Canales

package main

import (
	"fmt"
	"sync"
)

// le pasamos un array, un channel y el waitgroup
func sum(numbers []int, result chan int, wg *sync.WaitGroup) {
	defer wg.Done() //es lu ultimo que ejecuta la funcion.
	total := 0
	for _, number := range numbers { //sumamos los elementos del array
		total += number
	}
	result <- total //los seteamos en el channel
}

func main() {
	numbers1 := []int{1, 2, 3, 4, 5}
	numbers2 := []int{6, 7, 8, 9, 10}

	result1 := make(chan int) //creamos channel
	result2 := make(chan int) //creamos channel

	var wg sync.WaitGroup //iniciamos waitgroup
	wg.Add(2)             //le decimos q seran 2 gorutines

	go sum(numbers1, result1, &wg)
	go sum(numbers2, result2, &wg)

	go func() { //func anonima
		wg.Wait()
		close(result1)
		close(result2)
	}()

	total1 := <-result1
	total2 := <-result2

	fmt.Printf("Total1: %d, Total2: %d, Combined: %d\n", total1, total2, total1+total2)
}
