package main

import (
	"errors"
	"fmt"
)

// 1. Funciones Básicas
// Una función simple que no recibe parámetros y no devuelve valores.

func saludar() {
    fmt.Println("Hola, mundo")
}

// 2. Funciones con Parámetros
// Una función que recibe uno o más parámetros.

func saludarPersona(nombre string) {
    fmt.Println("Hola,", nombre)
}

// 3. Funciones que Devuelven Valores
// Una función que devuelve un valor. Go permite devolver múltiples valores desde una función.

func sumar(a int, b int) int {
    return a + b
}

// Devolviendo múltiples valores
func dividir(a float64, b float64) (float64, error) {
    if b == 0.0 {
        return 0.0, errors.New("división por cero")
    }
    return a / b, nil
}

// 4. Funciones Variádicas
// Una función que acepta un número variable de argumentos del mismo tipo.

func sumarVarios(numeros ...int) int {
    total := 0
    for _, num := range numeros {
        total += num
    }
    return total
}

// 6. Funciones como Valores y Cierres (Closures)
// Las funciones en Go pueden ser asignadas a variables y pasadas como argumentos a otras funciones. Los cierres son una especialidad de las funciones anónimas que pueden capturar variables fuera de su cuerpo.

func secuencia() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

// 7. Métodos
// En Go, un método es una función asociada a un tipo específico. Esto permite a la función "método" operar sobre los datos contenidos en el tipo de dato al que está asociada.

type Persona struct {
    nombre string
}

func (p Persona) saludar() {
    fmt.Println("Hola, mi nombre es", p.nombre)
}

// 8. Funciones con Parámetros por Nombre y Parámetros Opcionales
// Go no soporta directamente los parámetros por nombre o parámetros opcionales como tal, pero se pueden simular usando estructuras y campos con valores por defecto.

type Opciones struct {
    incremento int
    tieneSaludo bool
}

func avanzar(base int, opts Opciones) int {
    if opts.tieneSaludo {
        fmt.Println("Hola!")
    }
    return base + opts.incremento
}

// 9. Funciones con Valores de Retorno Nombrados
// Go permite nombrar los valores de retorno de las funciones, lo que puede hacer que el código sea más claro en su propósito.

func dividirConNombre(a, b float64) (resultado float64, err error) {
	if b == 0.0 {
		err = errors.New("división por cero")
		return
	}
	resultado = a / b
	return
}

func main() {

	saludar()

	saludarPersona("arnold")

	suma := sumar(3,2)
	fmt.Println(suma)

	division, _ := dividir(6,2)
	fmt.Println(division)

	suma_varios := sumarVarios(1,3,4,5,6)
	fmt.Println(suma_varios)

	// 5. Funciones Anónimas
	// Una función sin nombre, a menudo utilizada de manera inmediata o asignada a una variable.

	func() {
		fmt.Println("Función anónima")
	}()

	suma_anonima := func(a int, b int) int {
		return a + b
	}

	r_suma_anonima := suma_anonima(2,3)
	fmt.Println(r_suma_anonima)

	nextInt := secuencia()

	fmt.Println(secuencia()())

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	persona := Persona{nombre: "Arnold"}
	persona.saludar()

	avanzar_valor := avanzar(0, Opciones{incremento: 1, tieneSaludo: true})
	fmt.Println(avanzar_valor)

	division, _ = dividirConNombre(6, 3)
	fmt.Println(division)
}