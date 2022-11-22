package main

import "fmt"

func main() {

	// Diziler
	// dynamic -> dinamik
	// static  -> statik
	// 1,2,3
	// 0,1,2 -> index'ler 0'dan başlar

	// int , bool , float , string

	//var sayi int
	//var sayilar [3]int

	var siraNo int = 2
	//var yetkiliKisi string = "Hüsamettin"
	var girisler [3]string = [3]string{"Hüsamettin", "Uğur", "Ahmet"}

	yetkiKontrol(siraNo, girisler)

	//fmt.Println(yetkiliKisi)
	fmt.Println(girisler)

	/*var aktifKullanici string = girisler[siraNo]

	if aktifKullanici == yetkiliKisi {
		fmt.Println("Giriş Başarılı")
	} else {
		fmt.Println("Giriş Başarısız")
	}*/

	// index
	girisler[siraNo] = "Furkan"

	fmt.Println(girisler)

	//miktar := 107.50
	//miktarlar := [3]float64{107.50, 200.50, 300.50}

}

func yetkiKontrol(siraNo int, kullanicilar [3]string) {

	var yetkiliKisi string = "Ahmet"

	if kullanicilar[siraNo] == yetkiliKisi {
		fmt.Println("Giriş Başarılı")
	} else {
		fmt.Println("Giriş Başarısız")
	}

}
