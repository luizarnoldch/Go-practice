package main

import (
	"fmt"
	"math"
)

type Rectangulo struct {
	Ancho, Alto float64
}

func (r Rectangulo) Area() float64 {
	return r.Ancho * r.Alto
}

func (r Rectangulo) Perimetro() float64 {
	return 2 * (r.Ancho + r.Alto)
}

type Punto struct {
	X, Y float64
}

type Circulo struct {
	Centro Punto
	Radio  float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}

type Forma interface {
	Area() float64
}

func ImprimirArea(f Forma) {
	fmt.Println("Área de la forma:", f.Area())
}

func main() {
	rectangulo := Rectangulo{Ancho: 5, Alto: 5}
	circulo := Circulo{Centro: Punto{5,5}, Radio: 5}

	fmt.Println("Rectangulo")
	ImprimirArea(rectangulo)
	
	fmt.Println("Circulo")
	ImprimirArea(circulo)
}