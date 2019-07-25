package main

import ("fmt")

type car struct {
		Name string
		Model string
		Color string
		WeightInKg  int
}

func main () {
	c := car {
		Name : 		"Ferrari",
		Model: 		"GTC4",
		Color:		"Red",
		WeightInKg: 1920,
	}
	fmt.Println("Mobil : ", c)
}