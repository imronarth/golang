package main
import "fmt"

func main() {
	type NoKTP string

	var ktpImron NoKTP = "12321132312312"
	fmt.Println(ktpImron)
	fmt.Println(NoKTP("adsa"))
}