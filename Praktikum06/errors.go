package main

import (
	"errors"
	"fmt"
)
	
var (
	ValidationError = errors.New("Validayion error")
	NotFoundError = errors.New("not found error")
)

func GetById(id string) error{
	if id == "" {
		return ValidationError
	}

	if id != "Farah"{
		return NotFoundError
	}

	// Sukses

	return nil
}

func main(){
	err := GetById("Farah")
	if err != nil{
		if errors.Is(err, ValidationError){
			fmt.Println ("Validation error")
		}else if errors.Is(err, NotFoundError){
			fmt.Println("not found error")
		}else {
			fmt.Println("unknown error")
		}
	}
}
