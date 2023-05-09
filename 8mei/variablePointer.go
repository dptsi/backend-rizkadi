package main

import "fmt"

func main() {
    var x int = 10   // deklarasi variabel x dengan nilai 10
    var p *int = &x  // deklarasi variable pointer p dengan alamat memori dari x

    fmt.Println("Nilai x:", x)   // output: Nilai x: 10
    fmt.Println("Alamat x:", &x) // output: Alamat x: 0xc00000a0b8
    fmt.Println("Nilai p:", *p)  // output: Nilai p: 10
    fmt.Println("Alamat p:", p)  // output: Alamat p: 0xc00000a0b8

    *p = 20          // ubah nilai pada alamat memori yang ditunjuk oleh p menjadi 20
    fmt.Println("Nilai x setelah diubah:", x) // output: Nilai x setelah diubah: 20
}
