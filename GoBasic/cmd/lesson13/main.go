package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	//rand.Seed(time.Now().Unix())
	//rand.Intn(10) [0,n)
	//time.Sleep(time.Second)
	//fmt.Scanf()

	rastgeleSayi1 := rand.Intn(150)
	fmt.Println(rastgeleSayi1)

	time.Sleep(time.Second * 1)

	rastgeleSayi2 := rand.Intn(150)
	fmt.Println(rastgeleSayi2)

	fmt.Println("Sayı Giriniz : ")
	girilenSayi := 0
	fmt.Scanf("%d", &girilenSayi)
	// %d
	// %f
	// %s

	fmt.Println(girilenSayi)

	fmt.Println("İsimi Giriniz : ")
	girilenIsim := ""
	fmt.Scanf("%s", &girilenIsim)

	fmt.Println(girilenIsim)

	fmt.Println("Sayilar Giriniz : ")
	sayi1 := 0
	sayi2 := 0
	fmt.Scanf("%d %d", &sayi1, &sayi2)
	fmt.Println(sayi1)
	fmt.Println(sayi2)

	var sayilar [5]int = [5]int{0, 0, 0, 0, 0}
	fmt.Println("Kart Sayılarını Giriniz : ")
	fmt.Scanf("%d %d %d %d %d", &sayilar[0], &sayilar[1], &sayilar[2], &sayilar[3], &sayilar[4])
	fmt.Println(sayilar)

}
