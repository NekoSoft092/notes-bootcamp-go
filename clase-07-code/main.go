package main

import (
	"fmt"
	"time"
)

// Clase 1 del dia

func NewPerson(name string, age int) (person Person) {
	if name == "" && age == 0 {
		person = Person{
			Name: "Nel Alejandro",
			Age:  23,
		}
	} else {
		person = Person{
			Name: name,
			Age:  age,
		}
	}
	return
}

type Person struct {
	Name string
	Age  int
}

// Clase 2 del dia

// Repaso

// Composicion: Poder crear programas mas grandes a partir de piezas mas pequeñas, esto nos ayuda a
// diseñar distintos tipos de datos

// Ejemplo de embeddings en las clases
type Author struct {
	FirstName string
}

type Book struct {
	Id            string
	BookAtributes // Embeddings
}

type BookAtributes struct {
	Title       string
	Description string
	Author      Author
}

// Punteros - apuntadores
// Memory address 0x1h32 en hexa

func Increase(v *int) {
	(*v)++ // * accedemos al valor
}

// Person con apuntador
type Person1 struct {
	Name *string
	Age  int
}

// Interfaces
// Def: Es una forma de definir metodos que deben ser utilizados pero sin definirlos (Como un contrato), serie de funcionalidades
// Para que se utilizan: Brindan modularidad a la app, dleegacion de responsabilidades

// Funcion como interface -> Interfaz para una funcion
type PrepareFood func(ingredient string) bool

// Funcion que se adapta a la interfaz
func PrepareChicken(ingredient string) bool {
	return true
}

// Interfaces en struct
type Breakfast struct {
}
type Lunch struct {
}

type CookBreakfast func(ingredient string) Breakfast
type CookLunch func(ingredient string) Lunch

// Interface con struct, estructura que simula una interface
type Cook struct {
	Breakfast CookBreakfast
	Lunch     CookLunch
}

// Interface con interface
type ICook interface {
	CookBreakfast(ingredient string) Breakfast
	CookLunch(ingredient string) Lunch
}

// Cocinero impleneta ICook si implementa sus metodos
type Cocinero struct {
	Name  string
	Speed int
}

// Implemetar los metodos de cocinero para que impleemte ICook
func (c Cocinero) CookBreakfast(ingredient string) Breakfast {
	time.Sleep(time.Second * time.Duration(c.Speed))
	return Breakfast{}
}

func (c Cocinero) CookLunch(ingredient string) Lunch {
	time.Sleep(time.Second * time.Duration(c.Speed))
	return Lunch{}
}

func main() {

	// Clase 1 del dia

	// Clase 2 del dia
	author := Author{
		"Juan",
	}
	fmt.Println("El nombre del autor es ", author.FirstName)

	// Apuntadores
	var v int = 18
	Increase(&v) // & nos da la direccion en memoria de v (en hexa)
	fmt.Println("El valor de v es ahora ", v)

	var name string = "Nel Alejandro"
	var ptr *string = &name // string: 0xa31542

	var ptr2 *string = ptr // string: 0xa31542

	fmt.Println("Print ptr2 ", ptr2)

	// Shallow copy -> Copiamos una misma referencia de dos cosas distintas (al cambiar una cambia la otra tambien)
	p := Person1{Name: new(string), Age: 19}
	(*p.Name) = "Primer cambio"
	p1 := p
	(*p1.Name) = "Segundo cambio"

	fmt.Println("p.name: ", *p.Name)

	// Deep copy -> Copia en un nuevo apuntador (No hay misma referencia a dos elementos distinos "las dos copias")
	p2 := Person1{Name: new(string), Age: 20}
	(*p2.Name) = *p.Name
	(*p2.Name) = "Ultimo cambio"

	fmt.Println("p.name debe ser Segundo cambio", *p.Name)

	// Valor de un puntero por default es 0x0 = nil
	var punteroDefault *string
	fmt.Println("Valor por defecto de un puntero ", punteroDefault)

	//----------------
	// Interfaes

	var cocineroAlejandro ICook = Cocinero{
		Name:  "Nel Alejandro",
		Speed: 1,
	}

	breackfast := cocineroAlejandro.CookBreakfast("eggs")
	fmt.Println("yummy ", breackfast)

	// Interfaz vacia es como un tipo de dato any
	// Se utilizan

	json := map[string]any{
		"name": "Nel Alejandro Mora Pinto",
		"age":  23,
	}
	fmt.Println("json: ", json["name"])

}
