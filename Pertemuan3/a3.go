package main

import (
	"./hitungluas"
	"fmt")

	 //skema dipindahin ke sebelah
func main()  {
	var panjang float32 = 10.0
	var	lebar float32 = 5.0 

	_ , hasilkeliling := hitungluas.Hitungpersegipanjang(panjang, lebar)
	fmt.Println(hasilkeliling)
}