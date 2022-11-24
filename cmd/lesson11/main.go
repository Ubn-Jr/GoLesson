package main

import "fmt"

func main() {

	//var siraNo int = 0
	var geciciSayi int = 0
	var sayilar [5]int = [5]int{3, 7, 2, 11, 5}

	/*for i := 0; i < 5; i++ {

		geciciSayi = sayilar[i]
		geciciSayi *= geciciSayi
		fmt.Println(geciciSayi)

	}*/

	//continue -> Sonraki adıma geçer
	//break -> Döngüden çıkar
	/*for i := 0; i < 10; i++ {
		fmt.Println(i)
	}*/

	for i := 0; i < 5; i++ {
		geciciSayi = sayilar[i]
		if geciciSayi > 9 {
			continue
		}

		if geciciSayi%2 == 0 {
			break
		}

		geciciSayi *= geciciSayi
		fmt.Println(geciciSayi)
	}

	var rastgeleSayilar [25]int = [25]int{8, 43, 56, 786, 34, 23, 45, 67, 89, 12, 134, 56, 178, 90, 23, 145, 57, 89, 12, 34, 56, 78, 90, 23, 45}
	var tekSayiAdedi int = 0
	var ciftSayiAdedi int = 0
	var gecerliBirDiziMi bool = true

	for i := 0; i < len(rastgeleSayilar); i++ {

		var siradakiSayi int = rastgeleSayilar[i]
		fmt.Println(siradakiSayi)
		if siradakiSayi == 0 {
			gecerliBirDiziMi = false
			break
		}
		if siradakiSayi > 99 {
			continue
		}

		if siradakiSayi%2 == 0 {
			ciftSayiAdedi++
		} else {
			tekSayiAdedi++
		}
	}

	if gecerliBirDiziMi {
		fmt.Println("Tek Sayı Adedi: ", tekSayiAdedi)
		fmt.Println("Çift Sayı Adedi: ", ciftSayiAdedi)
	} else {
		fmt.Println("Dizi Geçersiz")
	}

	fmt.Println("Bitti")
	/*geciciSayi = sayilar[siraNo]
	geciciSayi *= geciciSayi
	fmt.Println(geciciSayi)
	siraNo++

	geciciSayi = sayilar[siraNo]
	geciciSayi *= geciciSayi
	fmt.Println(geciciSayi)
	siraNo++

	geciciSayi = sayilar[siraNo]
	geciciSayi *= geciciSayi
	fmt.Println(geciciSayi)
	siraNo++

	geciciSayi = sayilar[siraNo]
	geciciSayi *= geciciSayi
	fmt.Println(geciciSayi)
	siraNo++

	geciciSayi = sayilar[siraNo]
	geciciSayi *= geciciSayi
	fmt.Println(geciciSayi)*/

}
