package main

import "fmt"

func main() {
	// 0,1 -> bit bool
	// 00110101 -> byte

	// A -> 65 -> 01000001
	// h -> 104 -> 01101000
	// m -> 109 -> 01101101
	// e -> 101 -> 01100101
	// t -> 116 -> 01110100

	// 01000001 01101000 01101101 01100101 01110100

	// 512GB
	// 512 * 1024 MB
	// 512 * 1024 * 1024 KB
	// 512 * 1024 * 1024 * 1024 byte

	var sayi int = 2
	var sayi2 int = 6

	z := sayi * sayi2
	fmt.Println(z)

}
