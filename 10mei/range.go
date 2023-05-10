package main

import "fmt"

func main() {
	slice := []string{"apple", "banana", "orange"}
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}
	
	
}