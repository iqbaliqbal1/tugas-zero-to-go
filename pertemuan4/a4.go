package main 

import "fmt"

type Person struct {
	Name string
	age int 
}

func main()  {
	var numberA interface{}
	var numberB interface{}

	numberA = 10
	numberB = 20
	fmt.Println(numberA.(int) + numberB.(int))
}