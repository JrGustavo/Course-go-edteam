package main

import "fmt"

func main() {
	// operadores comparación:  >, <, ==, !=, <=
	fmt.Println(6 >= 4)

	// Operadores logiscos &&, //
	var age uint = 33
	fmt.Println("Es adulto?:", age >= 18 && age <= 60)
	fmt.Println("Es  niño o anciano?: ", age < 18 || age > 60)

	// Operador lógico Unario: !

	fmt.Println(!(4 != 4))
}
