package main

import "fmt"

type Ogrenci struct {
	Adi          string
	Yasi         int
	OkulAdi      string
	SinifNo      int
	Ortalama     float32
	vizeNotlari  []Not
	finalNotlari []Not
}

//Kalıtım
//OOP - Object Oriented Programming

type Not struct {
	Deger float32
}

func main() {

	// struct
	var ogrenciAdlari = []string{"Uğur", "Ahmet", "Mehmet", "Ali", "Veli"}
	var ogrenciYaslari = []int{20, 21, 22, 23, 24}

	var i int = 2

	fmt.Println(ogrenciAdlari[i], ogrenciYaslari[i])

	var ogr1 Ogrenci
	ogr1.Adi = "Uğur"
	ogr1.Yasi = 20
	ogr1.OkulAdi = "Yıldız Teknik Üniversitesi"
	ogr1.SinifNo = 1
	ogr1.Ortalama = 3.5

	fmt.Println(ogr1.OkulAdi)

	for {

		fmt.Scanf("%s\n", &ogr1.Adi)
		fmt.Scanf("%d\n", &ogr1.Yasi)
		fmt.Scanf("%s\n", &ogr1.OkulAdi)
		fmt.Scanf("%d\n", &ogr1.SinifNo)
		fmt.Scanf("%f\n", &ogr1.Ortalama)

		fmt.Println(ogr1)
	}

}
