package main

import "fmt"

type Hesaplama struct {
	Sayi1    float32
	Sayi2    float32
	Operator string
}

func (hesap Hesaplama) Hesapla() {
	if hesap.Operator == "+" {
		hesap.Toplama()
	} else if hesap.Operator == "-" {
		hesap.Cikarma()
	} else if hesap.Operator == "*" {
		hesap.Carpma()
	} else if hesap.Operator == "/" {
		hesap.Bolme()
	} else {
		fmt.Println("Geçersiz İşlem")
	}
}

func (hesap Hesaplama) Toplama() {
	var sonuc = hesap.Sayi1 + hesap.Sayi2
	fmt.Println(sonuc)
}

func (hesap Hesaplama) Cikarma() {
	var sonuc = hesap.Sayi1 - hesap.Sayi2
	fmt.Println(sonuc)
}

func (hesap Hesaplama) Carpma() {
	var sonuc = hesap.Sayi1 * hesap.Sayi2
	fmt.Println(sonuc)
}

func (hesap Hesaplama) Bolme() {
	var sonuc = hesap.Sayi1 / hesap.Sayi2
	fmt.Println(sonuc)
}

func main() {

	var hesap Hesaplama
	hesap.Sayi1 = 10
	hesap.Sayi2 = 20
	hesap.Operator = "+"

	hesap.Hesapla()

	var kare Kare
	kare.kenar = 10

	var dikdortgen Dikdortgen
	dikdortgen.en = 15
	dikdortgen.boy = 3

	fmt.Println(kare.AlanHesaplama())
	fmt.Println(dikdortgen.AlanHesaplama())

	var ogr1 Ogrenci
	ogr1.Adi = "Uğur Bilen"
	ogr1.Ortalama = 3.5

	var ogr2 Ogrenci
	ogr2.Adi = "Uğur Bilmez"
	ogr2.Ortalama = 1.5

	var ogrDrBel OgrenciDurumuBelirleyici
	ogrDrBel.ogr = ogr1
	ogrDrBel.DurumBelirle()

	ogrDrBel.ogr = ogr2
	ogrDrBel.DurumBelirle()

	//var ort Ogretmen

	var kumanda Kumanda
	kumanda.Ac()

}

type Ogrenci struct {
	Adi      string
	Ortalama float32
	Yasi     int
}

type Ogretmen struct {
	Adi  string
	Yasi int
}

type OgrenciDurumuBelirleyici struct {
	adet int
	ogr  Ogrenci
}

func (o OgrenciDurumuBelirleyici) DurumBelirle() {
	if o.ogr.Ortalama >= 2.5 {
		fmt.Println("Geçti")
	} else {
		fmt.Println("Kaldı")
	}
}

type Kare struct {
	kenar float32
}

func (k Kare) AlanHesaplama() float32 {
	return k.kenar * k.kenar
}

type Dikdortgen struct {
	en  float32
	boy float32
}

func (d Dikdortgen) AlanHesaplama() float32 {
	return d.en * d.boy
}

type Kumanda struct {
}

func (k Kumanda) Ac() {
	fmt.Println("Açıldı")
}

func (k Kumanda) Kapat() {
	fmt.Println("Kapatıldı")
}

func (k Kumanda) SesArttir() {
	fmt.Println("Ses Arttırıldı")
}

func (k Kumanda) SesAzalt() {
	fmt.Println("Ses Azaltıldı")
}
