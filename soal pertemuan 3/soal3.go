package main 

import "fmt"

func  generatepatterns(n int)  {
	for i := 1 ; i <= n ; i++{
		for j := 1 ; j <= i; j++{
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

func main () {

n := 3
generatepatterns(n)

n = 6
generatepatterns(n)
}