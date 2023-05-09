package main

import "fmt"

func main() {
	var point = 4
	switch point{
	case 1 :
		fmt.Println("perfect")
	case 2,3,4,5 :
		fmt.Println("almost")
	default :
		fmt.Println("wrong")
	}
}