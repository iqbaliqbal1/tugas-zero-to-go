package main

import ("fmt"
		"strings")

func main()  {
csv := map[string]string{}
var keys = "nama,nomor"
var values = "iqbal,122312"
splitteskeys := strings.Split(keys, ",")
fmt.Println(splitteskeys)
splittesvalues := strings.Split(values, ",")
fmt.Println(splittesvalues)
for i := 0; i <len(splitteskeys); i++{
	csv[splitteskeys[i]] =splittesvalues[i]
	}
	fmt.Println(csv)
}