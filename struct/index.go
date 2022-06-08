package main

import (
	"fmt"
	"os"
)

// type institute struct {
// 	instituteName          string
// 	instituteAddress       string
// 	instituteAccreditation string
// 	instituteSince         int
// }

// type course struct {
// 	name string
// }

// // Cara Buat Struct
// type student struct { // private access modifier: -> hanya bisa diakses di 1 package
// 	npm, name string    // variable/properti: -> private access modifier
// 	ins       institute // embedded struct atau casecade struct atau komposisi struct (spt inherintance)
// 	courses   []course
// }

// type highScoolStudent struct {
// 	nis, name string
// 	institute // embedded struct with syntatic sugar
// 	courses   []course
// }

// Tag Struct
// type Employee struct {
// 	FirstName string `json:"first_name"` // metadata informasi
// 	LastName  string `json:"last_name"`
// 	IsStatus  bool   `json:"is_status,omitempty"` // zero value -> (x)
// }

func main() {
	// 	// fmt.Println(student) // tidak bisa langsung memanggil struct tanpa di inisialisasi
	// 	var std student  // membuat sebuah variabel object dari struct dengan nama `std`
	// 	fmt.Println(std) // akan mencetak sebuah zero value dari msng" properti/variabel

	// 	// std.npm = "12345678" // isi value dari variabel/properti
	// 	// std.name = "Budi"
	// 	// std.grade = 100
	// 	// std.isDone = false

	// 	fmt.Println("std: ", std)
	// 	fmt.Println("std.name: ", std.name)

	// 	// menggunakan struct literal
	// 	var std2 = student{"87654321", "Joko", 90, true} // harus sama seperti yang di buat structnya
	// 	fmt.Println("std2: ", std2)
	// 	fmt.Println("std2.name: ", std2.name)

	// 	std3 := student{"828282929", "Sari", 80, false}
	// 	fmt.Println("std3: ", std3)

	// 	// deklarsi struct tanpa type alias langsung assign value
	// 	std4 := struct {
	// 		nama string
	// 	}{
	// 		"Jono",
	// 	}

	// 	fmt.Println("std4: ", std4)

	// 	// deklarsi struct tanpa type alias
	// 	var std5 struct {
	// 		alamat string
	// 	}

	// 	std5.alamat = "Ragunan"
	// 	fmt.Println("std5: ", std5.alamat)

	// 	// Pointer, di struct juga bisa kita terapkan
	// 	// mengambil sebuah reference dari std2
	// 	var stdPointer *student = &std2
	// 	fmt.Printf("std2 (address): %p\n", &std2)
	// 	fmt.Printf("stdPointer (address): %p\n", stdPointer)
	// 	fmt.Println("stdPointer.name: ", stdPointer.name)
	// 	// mengakses value yang bukan tipe data pointer, karena prop name: string
	// 	// fmt.Println("*stdPointer.name: ", *stdPointer.name)
	// 	fmt.Println("*stdPointer.name: ", (*stdPointer).name) // cara yang benar

	// 	// membuat sebuah variable dengan keyword new(), di struct juga bisa
	// 	// kita menggunakan student struct
	// 	var std3New = new(student) // mendefinisikan sebuah object struct menggunakan keyword new()
	// 	fmt.Println("std3New: ", std3New)
	// 	fmt.Println("std: ", std)

	// 	// tentang pembuatan yang menggunakan keyword make(): -> slices, maps, channels

	// xyzStudent := []student{
	// 	{
	// 		npm:  "13753030",
	// 		name: "Jution Candra Kirana",
	// 		ins: institute{
	// 			instituteName:          "Politeknik Negeri Lampung",
	// 			instituteAddress:       "Bandar Lampung",
	// 			instituteAccreditation: "B",
	// 			instituteSince:         1962,
	// 		},
	// 		courses: []course{
	// 			{
	// 				name: "Matematika Diskrit",
	// 			},
	// 			{
	// 				name: "Algoritma & Struktur Data",
	// 			},
	// 			{
	// 				name: "Dasar Pemrograman",
	// 			},
	// 		},
	// 	},
	// 	{
	// 		npm:  "13753031",
	// 		name: "Keysa",
	// 		ins: institute{
	// 			instituteName:          "Politeknik Negeri Lampung",
	// 			instituteAddress:       "Bandar Lampung",
	// 			instituteAccreditation: "B",
	// 			instituteSince:         1962,
	// 		},
	// 		courses: []course{
	// 			{
	// 				name: "Matematika Diskrit",
	// 			},
	// 			{
	// 				name: "Algoritma & Struktur Data",
	// 			},
	// 			{
	// 				name: "Dasar Pemrograman",
	// 			},
	// 		},
	// 	},
	// }

	// // fmt.Println(xyzStudent)
	// // fmt.Println()
	// // fmt.Printf("%+v \n", xyzStudent)

	// fmt.Println(strings.Repeat("-", 50))
	// for _, std := range xyzStudent {
	// 	fmt.Println(std.name, std.npm, std.ins.instituteName)
	// 	for _, course := range std.courses {
	// 		fmt.Println(course)
	// 	}
	// 	fmt.Println(strings.Repeat("-", 50))
	// }

	/*
				Challenge:
				-> Tambah Data (scanner)
				-> Lihat Data: 1. By Index, 2. All (scanner)
				-> Hapus Data By Index (scanner)

				Validasi:
				-> nama,npm, dll tidak boleh kosong
				-> index gk ada -> pesan
				repo: challenge-struct-student

			Soal

		Pada tugas ini kamu diminta untuk membuat sebuah aplikasi basis data mahasiswa sederhana berbasis console dengan kriteria yang sudah ditentukan

		1. Bisa menampung sampai dengan 5 mahasiswa atau lebih
		2. Validasi sesuai dengan screen, jika validasi tidak terpenuhi maka input tersebut akan terus diminta sampai terpenuhi
		3. Design screen harus mempunyai menu tambah mahasiswa, hapus mahasiswa, lihat mahasiswa dan exit
		4. Untuk menu tambah mahasiswa berikan validasi yaitu nama harus minimal 3 dan maksimal 20 karakter
		5. Untuk menu hapus mahasiswa akan otomatis index terakhir yang terhapus atau bisa by index
		6. untuk menu lihat mahasiswa akan ada sub menu lagi yaitu lihat menggunakan index dan lihat semua
		7. untuk menu lihat menggunakan index di input sesuai index data mahasiswa
		8. untuk menu lihat semua akan menampilkan seluruh data mahasiswa yang terdaftar
	*/

	// Masuk ke PART WEB API
	//DI Struct kita bisa memberikan sebuah tag (penanda)
	// Tapi biasanya case yang sering ditemukan ketika kita membuat sebuah data JSON, XML, CSV, YAML

	// emp := Employee{
	// 	FirstName: "Jack",
	// 	LastName:  "Mama",
	// }

	// // cetak format dari struct ke JSON
	// output, err := json.MarshalIndent(emp, "", " ")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(output))

	// env ? environment -> KEY-VALUE
	// Tiap OS bisa berbeda untuk cara set (mengisi si env secara dinamis) dan get (mengambil yang sudah diset)
	// GET -> echo %ENV_NAME% || echo %ENV_NAME (windows) | Terminal: command prompt, powershell (x)
	// GET -> echo $ENV_NAME (OsX, Linux)

	// SET -> set FULL_NAME=siapananda [enter] -> windows
	// SET -> FULL_NAME=siapaanda -> LINUX

	// Di golang sudah disediakan bagaimana kita mengkonsumsi sebuah env
	// Inbuilt OS Package

	// os.Setenv("FULL_NAME", "jution")   // SET Env
	fullName := os.Getenv("FULL_NAME") // GET Env
	fmt.Println(fullName)
}
