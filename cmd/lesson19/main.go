package main

// go build main.go

import (
	"errors"
	"fmt"
	"strings"
)

var icecekler = []string{"AYRAN", "SU", "KOLA"}
var yanUrunler = []string{"PATATES", "TATLI"}
var tostCesitleri = []string{"SADE", "KARISIK"}
var pizzaCesitleri = []string{"SUCUKSEVER", "MARGARITA"}
var burgerCesitleri = []string{"BIGMAC", "MCCHICKEN"}
var siparisler = []string{"TOST", "PIZZA", "BURGER"}
var fiyatlar map[string]float64 = map[string]float64{
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

	// Sipariş Seçimi
	_, siparisAdi, err := listedenElemanAlma("Sipariş", siparisler)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Ürün Seçimi
	_, urunAdi, err := urunCesitiSecimi(siparisAdi)
	if err != nil {
		fmt.Println(err)
		return
	}

	// İçeçek Seçimi
	_, icecekAdi, err := listedenElemanAlma("İçecek", icecekler)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Yan Ürün Seçimi
	_, yanUrunAdi, err := listedenElemanAlma("Yan Ürün", yanUrunler)
	if err != nil {
		fmt.Println(err)
		return
	}

	urunFiyati := fiyatlar[urunAdi]
	yanUrunFiyati := fiyatlar[yanUrunAdi]
	icecekFiyati := fiyatlar[icecekAdi]
	toplamFiyat := urunFiyati + yanUrunFiyati + icecekFiyati
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("Seçtiğiniz Sipariş : ", siparisAdi)
	fmt.Println("Seçtiğiniz Ürün : ", urunAdi)
	fmt.Println("Seçtiğiniz İçecek : ", icecekAdi)
	fmt.Println("Seçtiğiniz Yan Ürün : ", yanUrunAdi)
	fmt.Println("Ödemeniz Gereken Tutar : ", toplamFiyat)
	fmt.Println(strings.Repeat("=", 50))

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

	var secimYapilacakCesitler []string
	if siparisAdi == "TOST" {
		secimYapilacakCesitler = tostCesitleri
	} else if siparisAdi == "PIZZA" {
		secimYapilacakCesitler = pizzaCesitleri
	} else if siparisAdi == "BURGER" {
		secimYapilacakCesitler = burgerCesitleri
	} else {
		err = errors.New("Hatalı Giriş Yaptınız")
		return 0, "", err
	}

	urunNo, urunAdi, err := listedenElemanAlma("Ürün", secimYapilacakCesitler)

	return urunNo, urunAdi, err

}

func listedenElemanAlma(mesajUrun string, secimYapilacakUrunler []string) (int, string, error) {
	var err error
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("Lütfen Bir " + mesajUrun + " Seçiniz :")
	fmt.Println(-1, "İPTAL")
	for i := 0; i < len(secimYapilacakUrunler); i++ {
		fmt.Println(i, secimYapilacakUrunler[i])
	}

	var urunNo int
	var urunAdi string = ""
	fmt.Scanln(&urunNo)
	if urunNo == -1 {
		err = errors.New("Çıkış Yapıldı")
		return 0, "", err
	} else if (urunNo < 0) || (urunNo >= len(secimYapilacakUrunler)) {
		//err = errors.New("Hatalı Giriş Yaptınız")
		//return 0, "", err

		//RECURSİVE FONKSİYON - ASLA KULLANMA
		urunNo, urunAdi, err = listedenElemanAlma(mesajUrun, secimYapilacakUrunler)

	} else {
		urunAdi = secimYapilacakUrunler[urunNo]
		fmt.Println("Seçtiğiniz "+mesajUrun+" : ", urunNo, urunAdi)
	}

	return urunNo, urunAdi, err
}
