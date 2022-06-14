package main

import "fmt"

type stack []string

func (s *stack) size() int {
	return len(*s)
}

func (s *stack) push(name string) {
	*s = append(*s, name)
}

func (s *stack) pop() string {
	sz := s.size()
	data := (*s)[sz-1]
	*s = (*s)[0 : sz-1]
	return data
}

func main() {
	s := make(stack, 0)
	s.push("53")
	s.push("27")
	s.push("92")
	fmt.Println("Toples:")
	fmt.Println(s)

	t := s.pop()
	fmt.Println("Pop: ", t)
	t = s.pop()
	fmt.Println("Pop: ", t)
	fmt.Println(s)
}
