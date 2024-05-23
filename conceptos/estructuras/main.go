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

	fmt.Printf("%+v", junior)
}
