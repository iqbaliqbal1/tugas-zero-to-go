package main 

import "fmt"

func palindrome (s string) (palindrome bool){
	var huruf = len(s)
	for i := 0; i<huruf/2; i++ {
		if s[i] == s[huruf -1 - i]{
			palindrome = true
		}else{
			palindrome = false
			
		}
	}
	return palindrome
	
}

func main() {

	fmt.Println(palindrome("katak"))
	fmt.Println(palindrome("kabck"))
	fmt.Println(palindrome("cvc"))
}