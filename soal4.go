package main

import "fmt"

func main ()  {
	var bilangan int = 7
	var totalfaktor int = 0
	for i :=1 ; i <=bilangan ; i++{
		if bilangan %i == 0{
			// fmt.Println(i)
			totalfaktor = totalfaktor + 1
		}
	}
	if totalfaktor == 2 {
		fmt.Println("prima")
		
		}else{
		fmt.Println("tidak prima")
	} 
	
}	