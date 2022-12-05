package main

import (
	"fmt"
)

var (
	karakterListesiTr [12]rune = [12]rune{'Ş', 'ş', 'Ğ', 'ğ', 'Ü', 'ü', 'İ', 'ı', 'Ö', 'ö', 'Ç', 'ç'}
	karakterListesiEn [12]rune = [12]rune{'S', 's', 'G', 'g', 'U', 'u', 'I', 'i', 'O', 'o', 'C', 'c'}
)

// rune
// for - range
// "Ahmet" -> 'A', 'h', 'm', 'e', 't'
// rune[]

func main() {

	fmt.Println("Bir Kelime Giriniz:")
	kullaniciVerisi := ""
	fmt.Scanf("%s", &kullaniciVerisi)
	geciciVeri := ""
	for _, veri := range kullaniciVerisi {
		geciciHarf := veri
		for j := 0; j < len(karakterListesiTr); j++ {
			if veri == karakterListesiTr[j] {
				geciciHarf = karakterListesiEn[j]
				break
			} else {

			}
		}
		geciciVeri += string(geciciHarf)
	}
	fmt.Printf("Düzenlenmiş Veri : %s\n", geciciVeri)
}
