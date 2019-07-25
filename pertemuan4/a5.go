package main

import 
	"fmt"


type Person struct	{
	Name string
	Gender bool
	Age int 
	Height int 
}
func (P Person) PanggilNama() string {
	return "Hallo, Nama saya " + P.Name
}


func main () {
	iqbal := Person{
		Height: 170,
		Gender: false,
		Name: "M.Iqbal",
		Age: 23,
	}
	fmt.Println(iqbal)

	var iqbool = Person{"M.iqbool", true, 160, 27 }
	fmt.Println(iqbool)

	iqbabul := new(Person)
	iqbabul.Name = "Iqbabul"
	iqbabul.Gender = true
	iqbabul.Age = 22
	iqbabul.Height = 155
	fmt.Println(iqbabul)

	fmt.Println(iqbal.PanggilNama())
}