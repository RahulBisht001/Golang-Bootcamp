package main

import "fmt"

func main() {
	// variables
	age := 22
	name := "Rahul"
	height := 5.6

	//* Println vs Printf vs Print
	fmt.Println("Age is :", age, "Name is :", name, "Height is :", height)
	fmt.Println("Learning Golang!")

	fmt.Printf("Age is %d, Name is %s, Height is %f", age, name, height)
	fmt.Println("Learning Golang!")
}
