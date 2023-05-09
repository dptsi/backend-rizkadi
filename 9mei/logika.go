package main

import "fmt"

func main(){

	var valueTrue = true && true
	fmt.Println(valueTrue)

	var valueFalse = false && false
	fmt.Println(valueFalse)

	var valueFalse1 = true || false
	fmt.Println(valueFalse1)
}


