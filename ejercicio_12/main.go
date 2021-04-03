package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	readFileOne()
	readFileTwo()
	saveFileOne()
	saveFileTwo()
}

/**
Lee el archivo, pero no deja manipularlo.
**/
func readFileOne() {
	file, err := ioutil.ReadFile("./archivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		fmt.Println(string(file))
	}
}

/**
Lee y manipulacion del archivo
**/
func readFileTwo() {
	file, err := os.Open("./archivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		// lee y almacena en el buffer en memoria
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			// leer la linea
			registry := scanner.Text()
			// manipulacion del archivo
			fmt.Printf("Linea > %s \n", registry)
		}
	}
	file.Close()
}

/*
Grabar en archivo
*/
func saveFileOne() {
	file, err := os.Create("./newArchivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
		return
	}
	fmt.Fprintln(file, "Esta es una linea.")
	file.Close()
}

/*
Grabar en archivo
*/
func saveFileTwo() {
	fileName := "./newArchivo.txt"
	if !Append(fileName, "\nEsta es una segunda linea.") {
		fmt.Println("Error en la 2da linea.")
	}
}

func Append(fileName string, text string) bool {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}
	// cuando usamos "_" es por que no nos interesa grabar o leer el dato para la variable que devuelve la funcion.
	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}
	return true
}
