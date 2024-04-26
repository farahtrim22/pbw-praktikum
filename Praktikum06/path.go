package main

import (
	"fmt"
	"os"
)

func main() {
	var err error
	fileInfo, err := os.Stat("Farah")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if fileInfo.IsDir() {
		fmt.Println("Farah adalah sebuah direktori.")
	} else {
		fmt.Println("Farah adalah sebuah file.")
	}
}
