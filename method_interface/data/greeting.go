package data

import "fmt"

type Greeting struct { // Unexported (private)
	Name string
}

// type Greeting2 struct {
// 	Name    string
// 	Address string
// }

// func SayHi(name string) string { // Exported (public)
// 	return "Halo " + name
// }

// func SayHello() Greeting {
// 	return Greeting{
// 		Name: "Jack",
// 	}
// }

// Membuat sebuah method
func (g *Greeting) SayHi() *Greeting {
	fmt.Println("Halo saya", g.Name)
	fmt.Printf("g SayHi (address): %p \n", &g.Name)
	return g
}

func (g *Greeting) SayGoodBye(name string) {
	fmt.Printf("g SayGoodBye (address): %p \n", &g.Name)
	fmt.Println("Sampai jumpa lagi", name)
}

// func (g Greeting2) SayHi() {
// 	fmt.Printf("Halo saya %s, tinggal di %s", g.Name, g.Address)
// }
