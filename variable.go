package main
import "fmt"

func main() {
	var name string

	name = "Imron Athoriq Firjatulloh"
	fmt.Println(name)

	name = "Imron Athoriq"
	fmt.Println(name)
	
	var univ = "Universitas Brawijaya"
	fmt.Println(univ)
	univ = "Politeknik STMI Jakarta"
	fmt.Println(univ)
	
	// without var use ':=' to declare variable
	prodi := "Sistem Informasi"
	fmt.Println(prodi)
	prodi = "Instrumentasi"
	fmt.Println(prodi)
	
	// multiple declaration var
	var (
		firstName = "Imron"
		middleName = "Athoriq"
		lastName = "Firjatulloh"
	)
	fmt.Println(firstName, middleName, lastName)

}