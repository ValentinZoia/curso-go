package main

import "fmt"

func main() {
	// 	var buffer [256]byte

	// 	fmt.Println(len(buffer))

	// 	slice := buffer[10:20]
	// 	for i := 0; i < len(slice); i++ {
	// 		slice[i] = byte(i)
	// 	}
	// 	fmt.Println("before", slice)
	// 	AddOneToEachElement(slice)
	// 	fmt.Println("after", slice)

	// }

	// func AddOneToEachElement(slice []byte) {
	// 	for i := range slice {
	// 		slice[i]++
	// 	}

	//Array
	var array [4]int

	//puedo editar inices que quiero
	array[0] = 1
	array[1] = 2
	fmt.Println(len(array), cap(array)) // -->4 4

	//Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice)) // --> [0 1 2 3 4 5 6] 7 7

	//Agregando elementos al Slice
	slice = append(slice, 7)
	fmt.Println(slice, len(slice), cap(slice)) // --> [0 1 2 3 4 5 6 7] 8 14
	//Agregando otro slice al slice
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice, len(slice), cap(slice)) // --> [0 1 2 3 4 5 6 7 8 9 10] 11 14

	//Recorriendo Slice con range
	slice2 := []string{"hola", "que", "hace"}

	for i, valor := range slice2 {
		fmt.Println(i, valor)
		/*
			0 hola
			1 que
			2 hace
		*/
	}

}
