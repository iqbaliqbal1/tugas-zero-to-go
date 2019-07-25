package main 

import "fmt"

type Customer struct {
	Id int 
	Name string
	addr address
	married bool 
}

type address struct{
	Number int 
	City string
	Country string
}

func main(){
	var newCustomer = Customer{
		Id:1,
		Name: "John",
		addr: address{
			Number: 20 ,
			City: "Jakarta",
			Country: "Indonesia",
		},
		married : true,
	}
	fmt.Println("Hello, my name is", newCustomer.Name)
	fmt.Println("I live in", newCustomer.addr.City,newCustomer.addr.Country,"on jalan setiabudi no2b")
}