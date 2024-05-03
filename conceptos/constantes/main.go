package main

import "fmt"

const (
	os     = "linux"
	domain = "Team"
)

// Estas son las constantes de los meses
const (
	Jan = iota + 1
	Feb
	Mar
	Abr
	May
	Jun
	Jul
	Aug
	Sept
	Oct
	Nov
	Dec
)

func main() {
	fmt.Println(os, domain)
}
