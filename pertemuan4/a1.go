package main 

import "fmt"

func main() {

	var numberA int = 4
	var numberB *int = &numberA

fmt.Println(*numberB)
}