package data

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (s Student) ChangeNamePBV(name string, age int) {
	fmt.Println("name (changeNamePBV): ", name)
	s.Name = name
	fmt.Printf("s: \t%p \n", &s)
	fmt.Printf("s.Name: %p \n", &s.Name)
	fmt.Printf("s.Age: \t%p \n", &s.Age)
}

func (s *Student) ChangeNamePBR(name string, age int) {
	fmt.Println("name (changeNamePBR): ", name)
	fmt.Println("age (changeNamePBR): ", age)
	s.Name = name
	s.Age = age
	fmt.Printf("s: \t%p \n", s)
	fmt.Printf("s.Name: %p \n", &s.Name)
	fmt.Printf("s.Age: \t%p \n", &s.Age)
}
