package main

import "fmt"

func main()  {

	var total int 
	var arrayinput = [5]int{5,6,7,8,9}
	var ratarata int 
	for a := 0 ; a < len(arrayinput) ; a++ {
		total = total + arrayinput[a]
	}
	ratarata = total / len(arrayinput)
	fmt.Println(ratarata)
}