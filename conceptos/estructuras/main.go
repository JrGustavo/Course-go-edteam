package main

import "fmt"

type Person struct {
	Name        string
	Age         uint8
	HasChildren bool
}

func main() {
	junior := Person{
		Name:        "Junior",
		Age:         30,
		HasChildren: false,
	}

	//beto := Person{"Beto", 33, true}
	//javier := Person{Age: 32}

	alvaro := &junior
	(*alvaro).Name = "Alvaro"

	fmt.Printf("%+v", junior)
	//fmt.Printf("%+v", beto)
	//fmt.Printf("%+v", javier)
	fmt.Printf("%+v", alvaro)
}
