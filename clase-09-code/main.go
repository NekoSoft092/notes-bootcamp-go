package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Clase 1 de la clase
// Manejo de errores en go

func validateStatusCode(code int) error {
	if code > 299 {
		// Funcion New
		var err error = errors.New("this es an error 299 http") // Variable de error
		return err
	}
	return nil
}

// Error suntomizado
type CustomError struct {
	Msg  string
	Code int
}

func (e *CustomError) Error() string {
	return e.Msg + " " + strconv.Itoa(e.Code)
}

// Movies part
type Movies struct {
	Id     int
	Title  string
	Year   int
	Rating float64
}

// Clase 2 del dia
func panicFunction() {
	//  Def: Nos permite declarar un panic, podemos personalizarlo
	// dependiendo de las necesidades
	// panic("causa de panico") -> recibe como argumento un interface
	// Aborta la aplicacion

	fmt.Println("Comenzando...")
	_, err := os.Open("file1.txt")

	if err != nil {
		panic("Ha ocurrido un panic")
		// panic(err)
		fmt.Println("end")
	}

}

func main() {

	// Clase 1 del dia

	// Clase 2 del dia
	panicFunction()

}
