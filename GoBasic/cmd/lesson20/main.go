package main

import (
	"fmt"
)

func main() {
	// int -> Tamsayı
	// float -> Ondalıklı sayı
	// string -> Metin
	// bool -> evet, hayır (true, false)

	// if
	// else if
	// else

	// array -> dizi Aynı türden verileri tutar (sabit uzunluk)
	// slice -> dilim Aynı türden verileri tutar (değişken uzunluk)
	// map -> harita Aynı türden verileri tutar (anahtar değer)

	/*var sayilar []int = []int{121, 144, 169, 196, 225}
	for i := 0; i < len(sayilar); i++ {
		fmt.Println(sayilar[i])
	}*/

	// for -> döngü
	// error (string)

	//tanım
	//koşul
	//arttırma - eksiltme

	for i := 0; i <= 10; i++ {
		kare := i * i
		fmt.Println(kare)
	}
}

func topla(x int, y int) int {
	return x + y
}

func rastgeleSayiUret() int {
	return 1
}

func yazdirAdim(ad string) {
	fmt.Println("Hüsamettin ARABACI")
}

func yazdirIsim(isim string) {
	fmt.Println(isim)
}
