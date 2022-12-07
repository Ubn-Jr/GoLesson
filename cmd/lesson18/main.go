package main

import "fmt"

func main() {

	// Array isn't recommended to be used but it's good to know the characteristics of array in Go
	//slice
	//map

	/* Array */
	var urunlerA [3]string = [3]string{"Elma", "Armut", "Muz"}
	var adetlerA [3]int = [3]int{50, 10, 20}
	/* Slice */
	var urunlerS []string = []string{"Elma", "Armut", "Muz"}
	var adetlerS []int = []int{50, 10, 20}
	/* Map */ // Key - Value = Anahtar - Değer
	// Maplerin içinde sıralama önemli değil
	var urunlerM map[string]int = map[string]int{"Elma": 50, "Armut": 10, "Muz": 20}
	/* End */

	for i := 0; i < len(urunlerA); i++ {
		fmt.Println(urunlerA[i])
		fmt.Println(adetlerA[i])
	}

	for i, urun := range urunlerA {
		fmt.Println(urun)
		adet := adetlerA[i]
		fmt.Println(adet)
	}

	for i := 0; i < len(urunlerS); i++ {
		fmt.Println(urunlerS[i])
		fmt.Println(adetlerS[i])
	}

	for i, urun := range urunlerS {
		fmt.Println(urun)
		adet := adetlerS[i]
		fmt.Println(adet)
	}

	for urun, adet := range urunlerM {
		fmt.Println(urun)
		fmt.Println(adet)
	}

	yeniUrun := "Karpuz"
	yeniAdet := 500

	urunlerS = append(urunlerS, yeniUrun)
	adetlerS = append(adetlerS, yeniAdet)

	for i, urun := range urunlerS {
		fmt.Println(urun)
		adet := adetlerS[i]
		fmt.Println(adet)
	}

	fmt.Println("Karpuz eklendi")
	urunlerM["Karpuz"] = 500
	urunlerM["Elma"] = 100
	for urun, adet := range urunlerM {
		fmt.Println(urun)
		fmt.Println(adet)
	}

	delete(urunlerM, "Armut")

	fmt.Println("Armut silindi")
	for urun, adet := range urunlerM {
		fmt.Println(urun)
		fmt.Println(adet)
	}

	var yeniS []int = make([]int, 0) // nil
	var yeniM map[string]int = make(map[string]int)

	fmt.Println(yeniS)
	fmt.Println(yeniM)

	yeniS = append(yeniS, 10)
	yeniM["Elma"] = 50

	fmt.Println(yeniS)
	fmt.Println(yeniM)

	for i := 0; i < 3; i++ {
		gelenDeger := 0
		fmt.Scanf("%d", &gelenDeger)
		yeniS = append(yeniS, gelenDeger)
	}

	fmt.Println(yeniS)

}
