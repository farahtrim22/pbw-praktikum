package main

import (
	"fmt"
	"os"
)

func main() {
	var err error
	// Mengubah izin direktori
	err = os.Chmod("Farah", 0042)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Izin 'Farah' telah diubah menjadi 0042")
}
