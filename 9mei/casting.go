package main

import "fmt"

func main(){
		//konversi eksplisit
		var x int =10
		var y float64 = float64(x)
		fmt.Printf("x = %d, y = %f\n", x , y)

		//konversi implisit
		var a int = 10
		var b float64 = 2.5
		var c = float64(a) * b
		fmt.Printf("c = %f\n", c)
	
}