package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	name string
	age  int
}

type Student struct {
	arr []int
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	pers := Person{age: 18, name: "Duc"}
	//pers.name = "Duc"
	//pers.age = 18
	fmt.Println(pers, pers.name, pers.age)

	student := Student{arr: []int{1, 2, 3}}
	fmt.Println(student.arr)

	user := User{Name: "Duc", Age: 23}
	fmt.Println(user)
	byteArray, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray))
}
