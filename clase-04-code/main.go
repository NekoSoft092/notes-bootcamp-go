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

	namesSlice = append(namesSlice, "Pinto", "Elemento adicin")
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

//------------------
// Repaso
//------------------

// Arrays en go (De tamaño fijo)
func arraysExample() {
	// Def: Es una coleccion de items del mismo tipo de datos que se almacenan de forma contigua en la memoria
	// Def
	var a [2]string

	// asignacion de valores en el array
	a[0] = "Primer item"
	a[1] = "Segundo item"

	// Obtener valores
	fmt.Println("Primero: "+a[0], "Segundo: "+a[1])

}

// Slices en go (De tamaño variable)
func sliceExample() {
	// Def: Los slices nos permiten almacenar un conjunto de datos homogeneo
	// es decir, todos ellos del mismo tipo

	// Go se encarga del tamaño dinamicamente
	var s []string        // Declaracion del slice
	s = make([]string, 6) // Inicializacion del slice en memoria len(s)=6
	fmt.Println("El tamaño del slice es ", len(s))

	// Obtener un rango
	fmt.Println("Rango de los 4 primeros elementos ", s[1:4])

	// Obtener un valor del slice
	fmt.Println("Primer valor del slice ", s[0])

	// Propiedades de un slice
	// Longitud (len): El numero de elementos que contiene
	// Capacidad (cap): El numero de elementos del array subyacente, contando desde el primer elemento del segmento

	// Agregar a un slice
	s = append(s, "Nuevo elemento 1", "Nuevo elemento 2")

}

func mapExample() {
	// Def: Los maps nos permiten crear variables de tipo clave -> valor, definiendo un tipo de dato para las
	// claves y uno para los valores

	mymap1 := map[string]int{}     // Declaracion forma 1
	mymap2 := make(map[string]int) // Declaracion forma 2

	// La funcion de make nos sirve para inicializarlo pero no podremos introducir datos en la misma sentencia de inicializacion
	fmt.Println("El tamaño de mymap1 es ", len(mymap1))
	fmt.Println("El tamaño de mymap2 es ", len(mymap2))
}

// Main function
func main() {
	condicionales()
	switchFunc()
	switchFuncEspecific()
	switchFallthought()
	forFunc()
}
