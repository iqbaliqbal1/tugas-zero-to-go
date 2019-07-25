package main

import "fmt"

func main()  {
	func Hitungpersegipanjang(panjang float32, lebar float32)(float32,float32){
		luas := panjang *lebar
		keliling := 2*(panjang +lebar)
	
		return luas,keliling

}