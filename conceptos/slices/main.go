package main

import "fmt"

func main() {
	//Slice: Son apuntadores a un Array, no poseen datos

	things := [7]string{"🚗", "🚕", "🏎️", "🚔", "🚨", "🔔", "⏰"}
	cars := things[0:3]
	red := things[0:4]

	fmt.Println("Things:", things)
	fmt.Println("Cars:", cars)
	fmt.Println("Red:", red)

	fmt.Println("Cars[0]:", cars[0])

}
