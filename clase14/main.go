package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		sleepyStormtrooper(i)
	}
	time.Sleep(4 * time.Second)
}

func sleepyStormtrooper(id int) {
	time.Sleep(3 * time.Second)
	fmt.Printf("The Stormtrooper, number %d, is snoring\n", id)
}
