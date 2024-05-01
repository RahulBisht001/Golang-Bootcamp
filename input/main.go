package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Learning how to take input from the User
	fmt.Println("Enter your name: ")
	var Name string
	fmt.Scan(&Name)
	firstName := Name
	fmt.Println("My First name is : " + firstName)

	/*
		Scan method is good but the problem with scan method is that
		it read the only till white space.
		eg: input is :- Rahul Bisht
		so it will only read Rahul because its before first white space
	*/

	//* bufio method
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')

	fmt.Println(name)

}
