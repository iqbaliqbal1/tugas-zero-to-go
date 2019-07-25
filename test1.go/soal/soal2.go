package main  

import "fmt"

func main()  {
	var input = [...]float64{7,12,5,10,7,3,5,6}
	var totalangka int = len(input)
	var median float64 
	if totalangka %2 == 1 {
		median = input[totalangka / 2] 
	}else{
		median = (input[totalangka / 2] + input[totalangka / 2 - 1 ]) / 2
	}
fmt.Println(median)
}
