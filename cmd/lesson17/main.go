package main

import (
	"errors"
	"fmt"
)

var (
	sifiraBolmeHatasi = errors.New("Sıfıra Bölme Hatası")
)

func main() {
	//interface
	//OOP
	//tinyGO
	//Js Embed ***
	// Java Script ---
	x := 20
	y := 1
	sonuc, err := bolmeIslemi(x, y)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sonuc)
	}
}

func bolmeIslemi(sayi1 int, sayi2 int) (int, error) {
	if sayi2 == 0 {
		return 0, sifiraBolmeHatasi
	} else {
		sonuc := sayi1 / sayi2
		return sonuc, nil
	}
}
