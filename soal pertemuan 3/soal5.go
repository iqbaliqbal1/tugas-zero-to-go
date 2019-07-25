package main 

import "fmt"

func DeretGeometri(arr []int) (hasil bool) {
	var a,b int 
	a = len(arr)
	b = arr[1] / arr[0]
	for i := 0 ; i<a-1; i++{
		if arr[i+1] / arr[i] == b{
			hasil = true 
			}else{
				hasil = false
				break	
			}
		}
	return hasil
}

func main () {
	fmt.Println(DeretGeometri([]int{1, 3, 9, 27, 81,243}))
	fmt.Println(DeretGeometri([]int{1, 3, 9, 27, 80}))
}