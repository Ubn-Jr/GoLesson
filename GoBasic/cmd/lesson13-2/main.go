package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	kartSayisi  int = 10
	hamleSayisi int = 6
)

var (
	oyuncu1Kartlar [10]int = [10]int{4, 5, 0, 2, 8, 9, 7, 6, 3, 1}
	oyuncu2Kartlar [10]int = [10]int{6, 2, 9, 3, 1, 4, 5, 0, 7, 8}
)

func main() {
	rand.Seed(time.Now().Unix())
	ayrac := strings.Repeat("=", 50)

	fmt.Println("Oyun Başladı")
	fmt.Println(ayrac)
	//oyuncu1Hamle := rand.Intn(10)
	//oyuncu2Hamle := rand.Intn(10)
	oyuncu1Hamle := 0
	oyuncu2Hamle := 0
	fmt.Println("Oyuncu 1 Kart Seçimi :")
	fmt.Scanf("%d\n", &oyuncu1Hamle)
	fmt.Println("Oyuncu 2 Kart Seçimi :")
	fmt.Scanf("%d\n", &oyuncu2Hamle)

	fmt.Printf("Oyuncu 1 Atılan Zar : %d\n", oyuncu1Hamle)
	fmt.Printf("Oyuncu 2 Atılan Zar : %d\n", oyuncu2Hamle)
	fmt.Println(ayrac)
	oyunuKazanOyuncu := 0
	for i := 0; i < hamleSayisi; i++ {
		time.Sleep(time.Second)
		oyuncu1Hamle = oyuncu2Kartlar[oyuncu1Hamle]
		oyuncu2Hamle = oyuncu1Kartlar[oyuncu2Hamle]
		fmt.Printf("Oyuncu 1'in Seçtiği Kart : %d\n", oyuncu1Hamle)
		fmt.Printf("Oyuncu 2'in Seçtiği Kart : %d\n", oyuncu2Hamle)
		fmt.Println(ayrac)

		if oyuncu1Hamle == 0 && oyuncu2Hamle == 0 {
			oyunuKazanOyuncu = 0
			break
		}

		if oyuncu1Hamle == 0 {
			oyunuKazanOyuncu = 1
			break
		}

		if oyuncu2Hamle == 0 {
			oyunuKazanOyuncu = 2
			break
		}
	}

	if oyunuKazanOyuncu == 0 {

		if oyuncu1Hamle < oyuncu2Hamle {
			oyunuKazanOyuncu = 1
		} else if oyuncu2Hamle < oyuncu1Hamle {
			oyunuKazanOyuncu = 2
		} else {
			oyunuKazanOyuncu = 0
		}
	}

	fmt.Println(ayrac)
	fmt.Println("Oyun Bitti")
	if oyunuKazanOyuncu == 0 {
		fmt.Println("Oyun Berabere")
	} else {
		fmt.Printf("Oyunu %d. Oyuncu Kazandı\n", oyunuKazanOyuncu)
	}

}
