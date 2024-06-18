//Range, Close y Select en channels

package main

import "fmt"

func main() {
	c := make(chan string, 2)
	c <- "Mesaje1"
	c <- "Mensaje2"

	fmt.Println(len(c), cap(c)) // 2 2

	//Recorrido de channel
	//Range y close

	close(c) //cerrar channel

	for message := range c {
		fmt.Println(message) //Mesaje1  Mensaje2
	}

	//Otro ejemplo
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch) // si no lo pongo tira un error
	}()

	for n := range ch {
		fmt.Println(n)

	}

	//-----Select----

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "hello from ch1"
	}()

	go func() {
		ch2 <- "hello from ch2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received", msg2)
	}
	//se ejecuta el primero q este listo

	//-----Otro ejemplo------
	email1 := make(chan string)
	email2 := make(chan string)

	go mjs("Mensaje1", email1)
	go mjs("mensaje2", email2)

	select {
	case mjs1 := <-email1:
		fmt.Println("Email recibido de email1", mjs1)
	case mjs2 := <-email2:
		fmt.Println("Email recibido de email2", mjs2)
	}

}

func mjs(text string, c chan string) {
	c <- text
}

func doubleReturn(a int) (c, d int) { return a, a * 2 }
