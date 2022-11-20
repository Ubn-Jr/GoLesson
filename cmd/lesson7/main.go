package main

import "fmt"

func main() {
	//bool -> true , false
	var aktifMi bool = false

	//var sayi int = 10
	if aktifMi {

		fmt.Println("Çalışıyor")
	} else {

		fmt.Println("Duruyor")
	}

	var sayi int = 10

	if sayi >= 10 {

		fmt.Println("Sayı 10'dan büyük")

	} else {
		fmt.Println("Sayı 10'dan küçük")
	}

	var numara1 int = 10
	var numara2 int = 20

	//var buyukMu bool = numara1 > numara2
	//var kucukMu bool = numara1 < numara2
	//numara1 >= numara2
	//numara1 <= numara2
	//numara1 == numara2

	if numara1 > numara2 {
		fmt.Println("1 > 2")
	} else if numara1 == numara2 {
		fmt.Println("1 = 2")
	} else {
		fmt.Println("1 < 2")
	}

	var urun string = "AYAKKABI"
	// ==

	var odenecekTutar float32 = 0
	if urun == "UN" {
		odenecekTutar = 50
	} else if urun == "ÇAY" {
		odenecekTutar = 80
	} else if urun == "PEYNIR" {
		odenecekTutar = 120
	} else if urun == "EKMEK" {
		odenecekTutar = 4
	} else {
		odenecekTutar = 0
	}

	fmt.Println("Ödenecek Tutar :")
	fmt.Println(odenecekTutar)

	if odenecekTutar == 0 {
		fmt.Println("Ürün Hatalı")
	}

	var adet int = 50

	if adet > 10 {
		fmt.Println("Adet > 10")
	} else if adet > 20 {
		fmt.Println("Adet > 20")
	} else if adet > 30 {
		fmt.Println("Adet > 30")
	} else {
		fmt.Println("Hatalı")
	}

	var uretimAdedi int = 100
	var hammaddeAdedi int = 110
	var kritikAdet int = 5

	if hammaddeAdedi > uretimAdedi {

		if hammaddeAdedi-uretimAdedi > kritikAdet {
			fmt.Println("Üretim Mümkün")
		} else {
			fmt.Println("Üretim İmkansız - Kritik Stok")
		}

	} else {
		fmt.Println("Üretim İmkansız - Hammadde Yetersiz")
	}

}
