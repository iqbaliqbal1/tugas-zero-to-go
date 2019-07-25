package main

import "fmt"

func main ()  {
	var bilangan int = 12
	for i :=1 ; i <=bilangan ; i++{
		if bilangan %i == 0{
				fmt.Println(i)
			}
		
		}
	}