package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "imron"
	names[1] = "athoriq"
	names[2] = "firjatulloh"
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// deklarasi langusng
	// kalo kosong default terisi
	var values = [3]int{
		90,
		80,
	}
	fmt.Println(values)
	fmt.Println(len(values))

	// array tidak ditentukan berapa

	var arrrayMultiple = [...]int{
		20,
		21,
		22,
		23,
		24,
		25,
	}

	fmt.Println(arrrayMultiple)
	fmt.Println(len(arrrayMultiple))

	// disini tidak ada tipe untuk menghapus bisa di set default atau kosong
}
