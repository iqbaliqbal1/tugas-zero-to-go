package main

import "fmt"

func main ()  {

var Bilangan int = 7
	var hasil int = Bilangan %2
	if hasil == 0{
		fmt.Println("Genap")
			
	}else {
		fmt.Println("Ganjil")
	}
}