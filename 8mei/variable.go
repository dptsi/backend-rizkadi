package main

import "fmt"

func main() {
	const firstName string = "Rizky "
	var lastName string = "Adi"
	fmt.Print("halo ", firstName, lastName, "!\n")

	//karna constanta maka data tidak dapat di rubah
	// firstName = "Aku"
	lastName = "Harley"

	//tipe data numerik
	var bilanganBulat uint8 = 20
	//tipe data desimal 
	var bilanganDesimal = 2.12
	//tipe data boolean
	var varBool = true

	fmt.Println("perkenalkan ", firstName, lastName, "!\n")
	fmt.Println(bilanganBulat, bilanganDesimal, varBool)


}