package main

import "fmt"

func main()  {
	var listbuahsayur =[]string{"apel","jeruk","bayam","kangkung"}
	for index, nama := range listbuahsayur{
	fmt.Println("ini index ke ", index, "Datanya adalah ", nama )
	}
}