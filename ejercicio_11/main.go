package main

import "fmt"

type Humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
}

type Animal interface {
	respirar()
	comer()
	EsCarnivoro() bool
}

type Vegetal interface {
	ClasificacionVegetal() string
}

/* Genero Humano*/
type Hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
}

type Mujer struct {
	Hombre
}

func (h *Hombre) respirar() {
	h.respirando = true
}

func (h *Hombre) comer() {
	h.comiendo = true
}

func (h *Hombre) pensar() {
	h.pensando = true
}

func (h *Hombre) sexo() string {
	if h.esHombre {
		return "Hombre"
	} else {
		return "Mujer"
	}
}

func HumanoRespirando(human Humano) {
	human.respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando \n", human.sexo())
}

/*---------------------------------------------------*/
// Genero Animal
type Perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
}

func (p *Perro) respirar() {
	p.respirando = true
}

func (p *Perro) comer() {
	p.comiendo = true
}

func (p *Perro) EsCarnivoro() bool {
	return p.carnivoro
}

func AnimalesRespirar(an Animal) {
	an.respirar()
	fmt.Println("Soy un animal y estoy respirando")
}

func AnimalesCarnivoros(an Animal) int {
	if an.EsCarnivoro() {
		return 1
	} else {
		return 0
	}
}

func main() {

	// Pedro := new(Hombre)
	// Pedro.esHombre = true
	// HumanoRespirando(Pedro)

	// Maria := new(Mujer)
	// HumanoRespirando(Maria)

	totalCarnivoros := 0
	Dogo := new(Perro)
	Dogo.carnivoro = true
	AnimalesRespirar(Dogo)
	totalCarnivoros += AnimalesCarnivoros(Dogo)

	fmt.Printf("Total carnivoros %d", totalCarnivoros)
}
