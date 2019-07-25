package main

import "fmt"

func ucapanSelamat(hour int, nama string)  {

	if hour < 12 {
			fmt.Println("selamat pagi", nama)
		}else if hour < 19{
			fmt.Println("selamat sore", nama)
		}else {
			fmt.Println("selamat malam", nama)
	}

}

func main() {
	hour := 15
	nama := "iqbal"
	ucapanSelamat(hour, nama)

	jam := 7
	ucapanSelamat(jam, nama)
}