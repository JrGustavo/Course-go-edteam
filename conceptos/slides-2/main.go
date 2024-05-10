package main

import "fmt"

func main() {
	//len(): #De elementos en el slice
	//cap(): #De elementos del array origen, a partir del indice donde se creo el slice

	//animals := [6]string{"🦄", "🦅", "🦑", "🦆", "🦖", "🦧"}
	//pets := animals[1:3] //
	//pets = append(pets, "🦂", "🦉", "🦚")

	//pets[2] = "🦀"

	//fmt.Println("animals:		animals")
	//fmt.Println("pets:", pets)
	//fmt.Println("tamaño pets:", len(pets))
	//fmt.Println("capacidad pets:", cap(pets))

	pets := make([]string, 0, 3)
	pets = append(pets, "🦑", "🦆", "🦖")

	fmt.Println("pets:", pets)
	fmt.Println("tamaño pets:", len(pets))
	fmt.Println("capacidad pets:", cap(pets))

}
