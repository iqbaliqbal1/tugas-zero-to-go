package main

import "fmt"


func luassegitiga(alas int, tinggi int) float32{
	luas := float32(alas * tinggi / 2)
	var keliling float32 = float32(alas+tinggi)
	return luas

}
func main()  {
alas := 10 
tinggi := 10 
fmt.Println(luassegitiga(alas, tinggi))	

}