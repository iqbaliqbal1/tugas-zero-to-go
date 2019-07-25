package main

import "fmt"

func main ()  {
	var bil int = 3
	
	var n int

	var k int
	
	for i := 1; i <= bil ; i++{
		for n := i; n <= bil-1; n++ {
			fmt.Print(" ")
		}
		for j := n; j <= i-1; j++ {
			if j%2 == 0 {
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		for k = 2; k <= i; k++ {
			if i%2 == 0 && k%2 == 0 || i == k || k%2 == 1 && i%2 == 1 {
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}