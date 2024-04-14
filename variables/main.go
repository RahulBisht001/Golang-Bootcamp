package main

import "fmt"

func main() {
	fmt.Println("variables in Golang")


	var name string = "Rahul"
	fmt.Println(name)

	var version = 2
	fmt.Println(version)
	version = 20
	fmt.Println(version)

	var money int64 = 9241414114109999
	fmt.Println(money + 1)

	var decision = true
	fmt.Println(decision)

	const pi = 3.14
	// pi = 34
    // the above line will give error as we are trying to assign a constant value
	// cannot assign to pi (neither addressable nor a map index 
	fmt.Println(pi)

	personIncome := 100000000
	fmt.Println(personIncome + 2)
	
}
