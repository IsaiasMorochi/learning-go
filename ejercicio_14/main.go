package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// use_DEFER()
	use_PANIC()
}

func use_DEFER() {
	file := "test_file.txt"
	// Open crea un buffer en memoria.
	f, err := os.Open(file)

	// Antes de irse por cualquier otro lado, se va cerrar el archivo primero.
	defer f.Close()

	if err != nil {
		fmt.Println("Error abriendo el archivo")
		os.Exit(1)
	}
}

func use_PANIC() {
	defer func() {
		record := recover()
		if record != nil {
			// Nos permite que ademas de mostrar hora y fecha, hace una salida de programa os.Exit(1)
			log.Fatalf("Ocurrio un error que genero panic\n %v", record)
		}
	}()

	message := 1
	if message == 1 {
		panic("Se encontro el valor de 1")
	}
}
