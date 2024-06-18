package main

import "fmt"

func main() {
	//Defer
	defer fmt.Println("me ejecuto a lo ultimo de todo")
	fmt.Println("hola")
	fmt.Println("mundo")

	for i := 0; i < 10; i++ {
		fmt.Println(i)

		//Continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		//break
		if i == 7 {
			fmt.Println("Break")
			break
		}
	}
}
