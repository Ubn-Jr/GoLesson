package main

import (
	"errors"
	"fmt"
)

var icecekler = []string{"AYRAN", "SU", "KOLA"}
var yanUrunler = []string{"PATATES", "TATLI"}
var tostCesitleri = []string{"SADE", "KARISIK"}
var pizzaCesitleri = []string{"SUCUKSEVER", "MARGARITA"}
var burgerCesitleri = []string{"BIGMAC", "MCCHICKEN"}
var tostMenu = []string{"TOST", "ICECEK", "YANURUN"}
var pizzaMenu = []string{"PIZZA", "ICECEK", "YANURUN"}
var burgerMenu = []string{"BURGER", "ICECEK", "YANURUN"}
var siparisler = []string{"TOST", "PIZZA", "BURGER"}
var fiyatlar map[string]float32 = map[string]float32{
	"AYRAN":      2.50,
	"SU":         1.50,
	"KOLA":       3.50,
	"PATATES":    2.00,
	"TATLI":      3.00,
	"SADE":       5.00,
	"KARISIK":    7.00,
	"SUCUKSEVER": 10.00,
	"MARGARITA":  12.00,
	"BIGMAC":     15.00,
	"MCCHICKEN":  13.00,
}

func main() {

	fmt.Println("Ubn-Jr Restoran Sipariş Programına Hoşgeldiniz")
	fmt.Println("Lütfen Siparişinizi Seçiniz :")
	for i := 0; i < len(siparisler); i++ {
		fmt.Println(i, siparisler[i])
	}
	var siparisNo int
	var siparisAdi string = ""
	fmt.Scanln(&siparisNo)
	if (siparisNo < 0) || (siparisNo >= len(siparisler)) {
		fmt.Println("Hatalı Giriş Yaptınız")
		return
	} else {
		siparisAdi = siparisler[siparisNo]
		fmt.Println("Siparişiniz : ", siparisAdi)
	}

	urunNo, urunAdi, err := urunCesitiSecimi(siparisAdi)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Seçtiğiniz Ürün : ", urunNo, urunAdi)

	// İÇECEKLER
	// 1. AYRAN
	// 2. SU
	// 3. KOLA

	// YAN ÜRÜNLER
	// 1. PATATES
	// 2. TATLI

	// TOST MENU, PIZZA MENU, BURGER MENU

	// 1. TOST MENU
	//  - 1.1.a. SADE
	//  - 1.1.b. KARISIK
	//  - 1.2. İÇECEK
	//  - 1.3. YAN ÜRÜN

	// 2. PIZZA MENU
	//  - 2.1.a. SUCUKSEVER
	//  - 2.1.b. MARGARITA
	//  - 2.2. İÇECEK
	//  - 2.3. YAN ÜRÜN

	// 3. BURGER MENU
	//  - 3.1.a. BIGMAC
	//  - 3.1.b. MCCHICKEN
	//  - 3.2. İÇECEK
	//  - 3.3. YAN ÜRÜN

	// ÖDEME TUTARI

}

// urunCesitiSecimi sipariş adına göre ürün seçimi yapar
//
// # Return Values
//
// urunNo : Seçilen ürünün numarası
//
// urunAdi : Seçilen ürünün adı
//
// err : Hata durumunda döndürülen değer
func urunCesitiSecimi(siparisAdi string) (int, string, error) {

	var err error

	var secimYapilacakCsitler []string
	if siparisAdi == "TOST" {
		secimYapilacakCsitler = tostCesitleri
	} else if siparisAdi == "PIZZA" {
		secimYapilacakCsitler = pizzaCesitleri
	} else if siparisAdi == "BURGER" {
		secimYapilacakCsitler = burgerCesitleri
	} else {
		err = errors.New("Hatalı Giriş Yaptınız")
		return 0, "", err
	}

	fmt.Println("Lütfen Bir Ürün Seçiniz :")
	for i := 0; i < len(secimYapilacakCsitler); i++ {
		fmt.Println(i, secimYapilacakCsitler[i])
	}

	var urunNo int
	var urunAdi string = ""
	fmt.Scanln(&urunNo)
	if (urunNo < 0) || (urunNo >= len(secimYapilacakCsitler)) {
		err = errors.New("Hatalı Giriş Yaptınız")
		return 0, "", err
	} else {
		urunAdi = secimYapilacakCsitler[urunNo]
	}

	return urunNo, urunAdi, err

}
