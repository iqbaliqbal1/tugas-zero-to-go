package main	

import "fmt"

func faktorbilangan(n int)  {
	
	for i := 1 ; i <= n ; i++{
	if n %i == 0 {
			fmt.Print(i," ")
		}
	}
	fmt.Println()
}

func main () {
	n := 6
	faktorbilangan(n)
	n = 12
	faktorbilangan(n)
	// n := 14
	// faktorbilangan(14)
	// n := 16
	// faktorbilangan(16)
	// n := 20
	// faktorbilangan(20)

}