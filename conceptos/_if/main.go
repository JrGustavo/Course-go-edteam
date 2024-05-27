package main

import "fmt"

func main() {
	character := "🦸🏽‍️"
	canSearch := true

	switch {
	case !canSearch:
		fmt.Println("No esta permitida la busqueda")
	case character == "🦸🏽‍️" || character == "🦸🏽‍":
		fmt.Println("es un superheroe")
	case character == "🧟" || character == "🦸🏽‍":
		fmt.Println("es un spervillano")
	default:
		fmt.Println("Es ujn personaje normal ")

	}
}
