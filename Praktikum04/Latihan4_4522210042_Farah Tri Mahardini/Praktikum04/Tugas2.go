package main

import "fmt"

// Mengurutkan array dalam satu kali perulangan
func main() {
	fmt.Println("Bubble Sort")
	var arrayNumber = [20]int{4, 2, 7, 1, 9, 5, 3, 8, 6, 10, 15, 12, 18, 11, 16, 14, 20, 13, 19, 17}
	for i := 0; i < len(arrayNumber); i++ {
		for j := 0; j < len(arrayNumber)-i-1; j++ {
			if arrayNumber[j] > arrayNumber[j+1] {
				arrayNumber[j], arrayNumber[j+1] = arrayNumber[j+1], arrayNumber[j]
			}
		}
	}

	fmt.Println("Setelah dilakukan pengurutan.")
	fmt.Println(arrayNumber)
}