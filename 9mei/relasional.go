package main 

import "fmt"

func main() {
	var value = 3
	var isEqual = (value == 2 )

	//printf digunakan untuk mencetak string termasuk format specifier seperti %d, %f, %s 
	fmt.Printf("nilai %d (%t)\n", value, isEqual )
}
