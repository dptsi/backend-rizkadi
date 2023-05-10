package main

import "fmt"

func main() {
	slice := make([]int, 3, 5)
	fmt.Println(cap(slice)) // Output: 5
	
	
}