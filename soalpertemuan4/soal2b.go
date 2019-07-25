package main 

import	( 
	"fmt"
)

type Student struct {
	RollNumber int 
	Name string
}
func main (){
	var ps interface
	ps := &s 
	fmt.Println(ps)
	fmt.Println(ps.Name)
	//..
	fmt.Println(ps)
}