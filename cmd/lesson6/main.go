package main

import "fmt"

func main() {
	var sistemGirisSaati string = "18"
	var sistemGirisDakikasi string = "10"
	//var sistemGirisZamani string = sistemGirisSaati + ":" + sistemGirisDakikasi
	var sistemGirisZamani string = saatDakikaBirlestir(sistemGirisSaati, sistemGirisDakikasi)
	fmt.Println(sistemGirisZamani)

	var randevuSaati string = "19"
	var randevuDakikasi string = "15"
	//var randevuZamani string = randevuSaati + ":" + randevuDakikasi
	var randevuZamani string = saatDakikaBirlestir(randevuSaati, randevuDakikasi)
	fmt.Println(randevuZamani)

	var islemSaati string = "18"
	var islemDakikasi string = "30"
	//var islemZamani string = islemSaati + ":" + islemDakikasi
	var islemZamani string = saatDakikaBirlestir(islemSaati, islemDakikasi)
	fmt.Println(islemZamani)

}

func saatDakikaBirlestir(saat string, dakika string) string {
	zaman := saat + ":" + dakika
	fmt.Println(zaman)
	return zaman
}
