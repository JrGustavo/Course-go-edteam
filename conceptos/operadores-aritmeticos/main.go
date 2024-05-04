package main

import "fmt"

func main() {
	// Operadores Aritmeticos

	var a = (2 + 3) * 2
	fmt.Println(a)

	// Operadores  de asignación
	var b int = 5
	b -= 2
	fmt.Println(b)

	// Declaración post-incremento y post-decremento
	// (No son duna expresión sino una declaración

	var c int = 6
	c++
	fmt.Println(c)
}
