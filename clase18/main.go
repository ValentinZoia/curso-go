//Ejemplo avanzado usando range, close y select

package main

import (
	"fmt"
	"sync"
	"time"
)

// Función que representa una tarea a procesar
func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		time.Sleep(time.Second) // Simula el tiempo de procesamiento
		results <- task * 2     // Simula el resultado del procesamiento
	}
}

func main() {
	const numWorkers = 3
	const numTasks = 5

	tasks := make(chan int, numTasks)
	results := make(chan int, numTasks)
	var wg sync.WaitGroup

	// Iniciar los trabajadores
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	// Enviar tareas a los trabajadores
	for i := 1; i <= numTasks; i++ {
		tasks <- i
	}
	close(tasks) // Cerrar el canal de tareas para indicar que no hay más tareas

	// Esperar a que todos los trabajadores terminen
	go func() {
		wg.Wait()
		close(results) // Cerrar el canal de resultados cuando todos los trabajadores hayan terminado
	}()

	// Recibir y mostrar los resultados
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}
