package main

import (
	"fmt"
)

func main() {

	// mapScore := make(map[string]int) // Pembuatan map dengan keyword make
	// fmt.Printf("%T \n", mapScore)
	// // isi value
	// // mapScore["key"] = value

	// mapScore["Anggar"] = 75
	// mapScore["Adise"] = 70
	// mapScore["Roi"] = 85
	// mapScore["Roi"] = 90 // assign new value

	// fmt.Println(mapScore["Roi"]) // Cetak map dengan key
	// fmt.Println(mapScore)
	// fmt.Println(mapScore["Anggar"])
	// fmt.Println(mapScore["anggar"])

	// fmt.Printf("%v \n", mapScore)
	// fmt.Printf("%#v \n", mapScore)

	// fmt.Println()

	// Pembuatan map
	var month = map[int]string{
		1: "Januari",
		2: "Februari",
		3: "Maret",
		4: "April",
		5: "Mei",
	}

	// cetak menggunakan for loop
	// for key, value := range month {
	// 	fmt.Println(key, value)
	// }

	// // Sort Map
	// keys := make([]int, 0, len(month))
	// for k := range month {
	// 	keys = append(keys, k)
	// }
	// sort.Ints(keys)
	// for _, v := range keys {
	// 	fmt.Println(v, month[v])
	// }

	fmt.Println("before deleted: ", len(month))

	// Map bisa kita hapus elemennya menggunakan key
	// Perintahnya bisa menggunakan keyword delete()
	delete(month, 5)
	fmt.Println("after deleted: ", len(month))

	fmt.Println(month[5])

	// Cek keberadaan key apakah ada atau tidak dalam satu map
	// ketika membuat sebuah map, map ini ada return value, optional value (bool)
	// value, isExist := month[4]
	// if isExist {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("Item is not exist")
	// }

	// Persingkat bisa seperti ini
	// if value, isExist := month[5]; isExist {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("Item is not exist!")
	// }

	// Kombinasi Map dan Slice
	var students = []map[string]string{
		{"ID": "S001", "Name": "Paula", "Address": "Jakarta"},
		{"ID": "S002", "Name": "Baim", "Address": "Jakarta"},
		{"ID": "S003", "Name": "Rafi", "Address": "Tanggerang"},
	}

	for _, student := range students {
		fmt.Println(student["Name"])
	}

	// Latihan
	/*
		Bagaimana cara buat map dalam map
		map[key0:map[key0:value key1:value]]
		Cetak: x["key0"]["key1"]
	*/
}
