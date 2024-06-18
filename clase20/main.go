package main

import "fmt"

func main() {
	x, y := doubleReturn(5)
	HOLAmundo := "jajaja"
	var r int
	fmt.Println(x, y, r, HOLAmundo)
}

func doubleReturn(a int) (c, d int) { return a, a * 2 }
