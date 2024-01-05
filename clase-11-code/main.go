package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Clase 1 del dia

func readFunc() {
	text := "lorem ipsum"
	reader := strings.NewReader(text)

	buf := make([]byte, 8) // Un buffer de 8bytes

	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break // Termina de leer la linea
			}
			fmt.Println(err)
			return
		}

		fmt.Println("Data read", string(buf[:n]))

	}
}

func writeFile() {
	wr := bufio.NewWriter(os.Stdout)
	text := []byte("lorem ipsum dolor sil awet") // Convertimos el string a bytes

	_, err := wr.Write(text) // White al file
	if err != nil {
		fmt.Println(err)
		return
	}

}

func readFile() {
	// con el paquete os nos permite funcionalidades para interactuar con el propio sistema operativo
	//
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	buf := make([]byte, 8)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println(err)
			return
		}

		fmt.Println(string(buf[:n]))

	}
}

func createFile() {
	// Creamos los archivos on unos permisos reestablecidos

	//  d or file- owner - group - other
	// 4 read-only, 6 read and write , 7 read, write and execute
	f, err := os.OpenFile("file1.txt", os.O_CREATE|os.O_WRONLY, 0644) // Owner tiene pemisos de lectura y escritura si el archivo no existe
	if err != nil {
		fmt.Println(err)
		return
	}

	// os.append("") sobre escribir un archivo
	f.Write([]byte("Heloo world"))
}

// Model to load info of csv
type User struct {
	Id   int
	Name string
	Age  int
}

func readCSV() {
	// Open file
	f, err := os.Open("data.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	rd := csv.NewReader(f)

	rd.Read() // skip header del csv

	for {
		record, err := rd.Read() // lee por linea
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println(err)
			return
		}
		// Serializar la info
		id, err := strconv.Atoi(record[0])
		name, err := strconv.Atoi(record[1])
		age, err := strconv.Atoi(record[2])

		user = User{
			Id:   id,
			Name: name,
			Age:  age,
		}

		// Mostrar en pantalla los
		fmt.Println(record)

	}
}

func main() {

	// Clase 1 del dia

	readFile()
	createFile()

	// Segunda implementacion del reader
	/*rd := bufio.NewReader(os.Stdin)
	for {

	}*/

}
