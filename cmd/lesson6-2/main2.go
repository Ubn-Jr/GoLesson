package main

import "fmt"

func main() {

	gun := "15"
	ay := "11"
	yil := "1988"
	saat := "05"
	dakika := "30"

	sonuc := tarihBirlestir(gun, ay, yil, saat, dakika)
	fmt.Println(sonuc)

}

func tarihBirlestir(gun string, ay string, yil string, saat string, dakika string) string {

	tarih := gunAyYilBirlestir(gun, ay, yil)
	fmt.Println(tarih)
	zaman := saatDakikaBirlestir(saat, dakika)
	fmt.Println(zaman)
	tarihZaman := tarih + " " + zaman
	return tarihZaman

}

func gunAyYilBirlestir(gun string, ay string, yil string) string {

	tarih := ay + "/" + gun + "/" + yil
	return tarih

}

func saatDakikaBirlestir(saat string, dakika string) string {
	zaman := saat + ":" + dakika
	return zaman
}
