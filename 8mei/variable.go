package main

import "fmt"

func main() {
	const firstName string = "Rizky "
	var lastName string = "Adi"
	fmt.Print("halo ", firstName, lastName, "!\n")

	//karna constanta maka data tidak dapat di rubah
	// firstName = "Aku"
	lastName = "Harley"

	fmt.Print("perkenalkan ", firstName, lastName, "!\n")
}