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
}

type Mujer struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
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
	return "Hombre"
}
func (m *Mujer) respirar() {
	m.respirando = true
}

func (m *Mujer) comer() {
	m.comiendo = true
}

func (m *Mujer) pensar() {
	m.pensando = true
}

func (m *Mujer) sexo() string {
	return "Mujer"
}

func HumanoRespirando(human Humano) {
	human.respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando \n", human.sexo())
}

func main() {

	Pedro := new(Hombre)
	HumanoRespirando(Pedro)

	Mujer := new(Mujer)
	HumanoRespirando(Mujer)
}
