package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) setName(name string) Person {
	if name == "" {
		return p
	}
	p.Name = name
	return p
}

func (p *Person) setAge(age int) {
	if age <= 0 {
		return
	}
	p.Age = age
}

func main() {
	person := Person{Name: "Duc", Age: 23}
	fmt.Println(person)
	person = person.setName("Vu")
	fmt.Println(person)

	(&person).setAge(20)
	fmt.Println(person)

}
