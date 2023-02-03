package main

import "fmt"

var allAnimals [30]string = [30]string{"Cat", "Dog", "Bird", "Fish", "Rabbit", "Horse", "Cow", "Pig", "Goat", "Sheep", "Chicken", "Duck", "Turkey", "Lion", "Tiger", "Bear", "Elephant", "Giraffe", "Zebra", "Monkey", "Gorilla", "Kangaroo", "Raccoon", "Squirrel", "Rat", "Mouse", "Deer", "Hippo", "Camel", "Panda"}
var allAnimalHandCounts [30]int = [30]int{4, 4, 2, 0, 4, 4, 4, 4, 4, 4, 2, 2, 2, 4, 4, 4, 4, 4, 4, 2, 2, 2, 4, 4, 4, 4, 4, 4, 4, 2}
var allAnimalTeethCounts [30]int = [30]int{30, 42, 0, 0, 28, 40, 32, 32, 32, 32, 0, 0, 0, 30, 30, 42, 0, 0, 0, 32, 32, 32, 28, 28, 28, 28, 0, 0, 0, 32}

func main() {

	var sayi1 int = 10
	var sayi2 int = 20
	//var toplam int
	//var fark int
	//var bolum int
	//var carpim int

	//toplam := toplamaIslemi(sayi1, sayi2)
	toplam, _, _, _, mesaj := dortIslem(sayi1, sayi2)

	fmt.Println(toplam)
	//fmt.Println(fark)
	//fmt.Println(bolum)
	//fmt.Println(carpim)
	fmt.Println(mesaj)

	uzunluk := len(allAnimalHandCounts)
	fmt.Println(uzunluk)
	for _, animal := range allAnimals {
		fmt.Println(animal)
	}

	for i := 0; i < len(allAnimals); i++ {
		fmt.Println(allAnimals[i])
	}

	var ornekKelime string = "Merheba Dünya" // " string
	//ornekKelime[4] = 'a' // ' char // ascii
	// char = 1 byte = 8 bit = 256 karakter
	// rune = 4 byte = 32 bit = 4 byte = 2^32 karakter

	karakter := ornekKelime[9]
	fmt.Println(string(karakter))

	fmt.Println(ornekKelime)

	for _, karakter := range ornekKelime {
		fmt.Println(string(karakter))
	}

}

func toplamaIslemi(sayi1 int, sayi2 int) int {
	return sayi1 + sayi2
}

func dortIslem(x int, y int) (int, int, int, int, string) {
	toplam := x + y
	fark := x - y
	bolum := x / y
	carpim := x * y

	mesaj := "Merhaba"

	return toplam, fark, bolum, carpim, mesaj
}
