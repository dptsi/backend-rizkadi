package main

import "fmt"

func main() {
	var fruits = [4]string{"apple","melom","orange"} //array berukuran 4 dengan data yang di isi hanya 3
	fmt.Println("jumlah buah \t\t",len(fruits)) // \t adalah tab
	fmt.Println("isi semua element \t", fruits)
	fmt.Println(fruits[0])// menampilkan data pertama dalam array
}