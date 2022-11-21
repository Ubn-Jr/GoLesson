package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	// && -> And Ve
	// || -> Or  Veya

	// true && true = true
	// false || false = false

	var turkce bool = true
	var math bool = true

	if turkce && math {
		fmt.Println("Tebrikler")
	} else {
		fmt.Println("Tekrar Dene")
	}

	var turkceNotu int = 80
	var matematikNotu int = 90
	var gecerNot int = 50

	//var turkceDurum bool = turkceNotu > gecerNot

	if turkceNotu > gecerNot && matematikNotu > gecerNot {
		fmt.Println("Tebrikler")
	} else {
		fmt.Println("Tekrar Dene")
	}

	var kilo int = 120
	var antreman int = 10

	//kilo > 100 && antreman > 8 // true
	//kiloe < 100 && antreman > 3 // true

	if (kilo > 100 && antreman > 8) ||
		(kilo < 100 && antreman > 3) {
		// Evet
		fmt.Println("Tebrikler")
	} else {
		// Hayir
		fmt.Println("Tekrar Dene")
	}

	// if else
	// > < >= <= ==
	// && ||

}
