package main

import (
	"fmt"
	"time"
)

func main() {
	// Canal
	channelOne := make(chan time.Duration)

	// GoRoutine
	go bucle(channelOne)
	fmt.Println("Llegué hasta aquí")

	// similar a un await de Node
	// El flujo de ejecución espera que channelOne tenga un valor para asignarlo a la variable.
	msg := <-channelOne
	fmt.Println(msg)
}

func bucle(channel chan time.Duration) {
	init := time.Now()
	for i := 0; i < 100_000_000_000; i++ {
	}
	final := time.Now()
	// al canal le asigna la duración - diferencia de init vs final.
	channel <- final.Sub(init)
}
