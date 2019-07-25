package main

import "fmt"

type person struct {
	Name string 
	Age int 
	Height int
	Gender bool 

}

func main (){
	 iqbal := person {
		Name : "iqbal iqbal",
		Age : 23,
		Height : 250,
		Gender : false,
	}
	fmt.Println(iqbal.Name)
}

func (p person) Panggilnama() string {
	return "asd" + p.Name
}{
	fmt.Println(iqbal.Panggilnama)
}
