package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Student struct {
	arr []int
}

func main() {
	pers := Person{age: 18, name: "Duc"}
	//pers.name = "Duc"
	//pers.age = 18
	fmt.Println(pers, pers.name, pers.age)

	student := Student{arr: []int{1, 2, 3}}
	fmt.Println(student.arr)
}
