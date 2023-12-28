package main

// GlobalVariable
var GlobalVariable string = "Global variable"

const (
	FuelType string = "constante e"
)

/*
 Esto es otro tipo de comentario de documentacion
*/

func main() {
	var fullname string = "Nel Alejandro Mora pinto"
	if fullname == "" {

	}

	// Estrucuras de control
	if ok := foo(); ok {

	}

	if n := 1; n == 1 {
		print("siempre entrara al if")
	}

}

func foo() bool {
	return true
}

var SumResult int = 0

func sum(num1 int, num2 int) {
	SumResult = num1 + num2
}
