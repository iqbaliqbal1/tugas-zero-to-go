package main	

import "fmt"

func prima(n int)  {
	var totalfaktor int 
	for i :=1 ; i <=n ; i++{
		if n %i == 0{
			totalfaktor = totalfaktor + 1
		}
	}
	if totalfaktor == 2 {
		fmt.Println("prima")
		
		}else{
		fmt.Println("tidak prima")
		}
	}

func main () {
	n := 7
	prima(n)
}