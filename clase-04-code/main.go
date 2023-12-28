package main

import (
	"fmt"
	"math/rand"
)

// Repaso

// Condicionales
func condicionales() {
	var salary float64 = 4000
	if salary <= 3000 {
		fmt.Println("Esta persona no debe pagar impuestos")
	} else if salary <= 4000 {
		fmt.Println("Esta persona debe pagar impuestos de 12%")
	} else {
		fmt.Println("Esta persona debe pagar impuestos de 15%")
	}
}

// Switch
func switchFunc() {
	var age uint8 = 18
	switch age >= 15 {
	case age >= 150:
		fmt.Println("Eres importal")
	case age >= 18:
		fmt.Println("Eres mayor de edad")
	default:
		fmt.Println("No eres mayor de edad aun")
	}
}

// Switch para un caso en especifico
func switchFuncEspecific() {
	day := "sunday"
	switch day {
	case "monday", "thuestday", "sunday":
		fmt.Println("Es un dya de la semana valido")
	default:
		fmt.Println("No es un dia de la semana valido")
	}
}

// Switch con fallthought
func switchFallthought() {
	i := 30

	switch {
	case i%5 == 0:
		fmt.Println("Es multiplo de 5")
		fallthrough
	case i%10 == 0:
		fmt.Println("Es umltiplo de 10") // Motrara que entro alas dos coondiciones, evalua las dos
	default:
		fmt.Println("no match")
	}
}

// for bucle
func forFunc() {
	// fluits := []string{"apple", "", ""}
	// time.Sleep(time.second)
	var counter int
	for counter <= 10 {
		fmt.Println("Primer bucle counter: ", 10)
		counter++
	}

	// Bucle infinito en el que tengamos una condicion de break
	var counter1 int
	for {
		fmt.Println("Segundo bucle counter1: ", counter1)
		if counter1 == 10 {
			break
		}
		counter1++
	}

	// For que usa range (limitado en un rango)
	for i := 0; i <= 10; i++ {
		fmt.Println("Este es un numero random ", rand.Int()%1000)
	}

	// Arrays
	apples := [2]string{"red", "green"}
	apples[0] = "yellow"
	fmt.Println("First apple", apples[0])

	// Array declarado con var
	var array [3]string
	fmt.Println("Array declarado con var", len(array))

	// Ingresar un valor en memoria dinamca (arraylist)
	// names := []string{} //no tiene un tamaño definido
	namesSlice := make([]string, 3, 3) // slice con tamaño inicial
	namesSlice[0] = "Nel"
	namesSlice[1] = "Alejandro"
	namesSlice[2] = "Mora" //cap(names) -> capacidad

	append(namesSlice, "Pinto", "Elemento adicin")
	fmt.Println("Nuevo array con append", namesSlice)

	// Maps key -> value (como un json)
	namesMaps := make(map[string]string)
	namesMaps["name"] = "Nel Alejandro"

	// Bucle infinito clasico
	/*
		sum:=0
		for {
			sum++
		}*/
}

func main() {
	condicionales()
	switchFunc()
	switchFuncEspecific()
	switchFallthought()
	forFunc()
}
