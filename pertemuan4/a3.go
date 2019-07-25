package main 

import "fmt"

type Person struct {
	Name string
	age int 
}

func main()  {
	var secret interface {}

	secret = "ethan hunt"
	fmt.Println(secret)

	secret = []string{"apple","manggo","banana"}
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)

	secret = Person {
		Name: "John",
		age: 20,
	}
	fmt.Println(secret)
}