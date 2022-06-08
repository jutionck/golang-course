package main

import "golang-fundamental/file/data"

var fileLocation = "/Users/jutioncandrakirana/Documents/GitHub/enigma/GOLANG/golang-fundamental/file/db.dat"

func main() {
	// data.CreateFile(fileLocation)
	// data.WriteToFile(fileLocation, "Ipin")
	data.OpenWithOsOpenFile(fileLocation)
}

/*
	1. Buat File (v)
	2. Membaca File (v)
	3. Menulis File (v)
	4. Menghapus File
*/
