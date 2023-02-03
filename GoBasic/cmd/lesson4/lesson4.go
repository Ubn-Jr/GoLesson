package main

import "fmt"

func main() {

	var sayi1 int = 10
	var sayi2 int = 20

	// + - * / %
	var toplam int
	toplam = sayi1 + sayi2
	fmt.Println(toplam)

	var cikarma int = sayi1 - sayi2
	fmt.Println(cikarma)

	carpma := sayi1 * sayi2
	fmt.Println(carpma)

	var bolme1 int = 0
	bolme1 = sayi2 / sayi1
	fmt.Println(bolme1)

	sayi3 := 11.0
	sayi4 := 4.0

	bolme2 := sayi3 / sayi4
	fmt.Println(bolme2)

	// 1,2,3,4,5,6,7,8,9
	// 1,2,3
	// 4,5,6
	// 7,8,9

	// Tek Çift %2

	// y = ax+b
	// y = 4a + b
	sayi5 := 5
	sayi6 := 4
	mod1 := sayi5 % sayi6
	fmt.Println(mod1)

	deger1 := 3
	deger2 := 2

	deger1 = deger1 + deger2
	deger1 += deger2

	deger1 = deger1 - deger2
	deger1 -= deger2

	deger1 = deger1 * deger2
	deger1 *= deger2

	deger1 = deger1 / deger2
	deger1 /= deger2

	deger1 = deger1 % deger2
	deger1 %= deger2

	numara1 := 12.7
	//numara2 := 1

	//numara1 = numara1 + numara2
	//numara1 += numara2

	fmt.Println(numara1)
	numara1++
	fmt.Println(numara1)
	numara1--

	var adim string = "Hüsamettin"
	var soyadim string = "ARABACI"
	var tamAdim string
	tamAdim = adim + " " + soyadim

	fmt.Println(tamAdim)

	var adim2 string = "Hüsamettin"
	var soyadim2 string = "ARABACI"

	adim2 += " "
	adim2 += soyadim2
	fmt.Println(adim2)

	// 0,1,2,3,4,5,6,7,8,9,10,11 -> Kalem
	// 0,2,4,6,8 -> 2
	// 1,3,5,7,9

	// Kutu Hacmi -> % 3
	// 0 % 3 = 0
	// 1 % 3 = 1
	// 2 % 3 = 2
	// 3 % 3 = 0
	// 4 % 3 = 1
	// 5 % 3 = 2
	// 6 % 3 = 0
	// 7 % 3 = 1
	// 8 % 3 = 2
	// 9 % 3 = 0

	// 0 1 2 -> Kutu1
	// 3 4 5 -> Kutu2
	// 6 7 8 -> Kutu3
	// 9 * * -> Kutu4

}
