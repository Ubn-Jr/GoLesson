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

}
