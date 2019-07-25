package main

import "fmt"

func main (){
	var a int = 3
    var k int 
    for i := 1; i <= a; i++ {     
        k=0
        for spasi := 1; spasi <= a-i; spasi++ {
            fmt.Print("  ")         
        }
        for {
            fmt.Print("*")
            k++
            if k == 2*i-1{             
                
            }
        }       
        fmt.Println("")
    }
}