package main
import "fmt"

func main() {
	var a = 10
	var b = 30
	var d = 5
	var e = 2
	var c = a + b - e * d
	var f = -12
	fmt.Println(f)

	aTrue := true
	fmt.Println("Hasil boolean a", aTrue)
	fmt.Println("Hasil boolean a not", !aTrue)

	fmt.Println("Hasil", c)
	
	a +=10
	fmt.Println("Hasil", a)
	
	a++
	fmt.Println("Hasil", a)
	
	fmt.Println("\nPerbandingan Operator")

	var name1 = "Imron"
	var name2 = "Imron"
	var name3 = "Athoriq"

	var boolVar bool = name1 == name2
	fmt.Println("Hasil Bool =", boolVar)
	boolVar = name1 != name3
	fmt.Println("Hasil Bool =", boolVar)
}