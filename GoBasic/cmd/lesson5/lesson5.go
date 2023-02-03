package main

import "fmt"

func main() {

	// x2y sonucunu tekra alıp x2y
	var x int = 2
	var y int = 3
	var z int
	z = x2y(x, y)
	fmt.Println(z)

}

func x2y(x int, y int) int {

	sonuc := (x * x) + y
	return sonuc
}

func simpleFunc() {
	//Body
}

func paramFunc(adi string, soyadi string, yasi int) {
	//Body
}

func returnFunc() int {

	return 0

}
