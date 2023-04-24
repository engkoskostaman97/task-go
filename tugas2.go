package main

import "fmt"

func main() {
    var angka int
    fmt.Print("Masukkan angka: ")
    fmt.Scanln(&angka)
    
    if angka%2 == 0 {
        fmt.Println(angka, "adalah bilangan genap")
    } else {
        fmt.Println(angka, "adalah bilangan ganjil")
    }
}
