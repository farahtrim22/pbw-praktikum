package main

import (
	"fmt"
)

func main() {
	var nama string
	var usia int

	fmt.Print("Masukkan Nama Anda: ")
	fmt.Scanln(&nama)

	fmt.Print("Masukkan Usia Anda: ")
	fmt.Scanln(&usia)

	kategori := "Dewasa"
	if usia < 18 {
		kategori = "Anak-Anak"
	}

	fmt.Printf("Selamat datang, %s Anda termasuk kategori %s.\n", nama, kategori)
}