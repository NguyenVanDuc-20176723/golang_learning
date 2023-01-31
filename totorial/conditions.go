package main

import "fmt"

func main() {
	a := 4
	switch a {
	case 1:
		fmt.Println("a = 1")
	case 2, 3:
		fmt.Println("a = 2,3")
	default:
		fmt.Println("a = ?")
	}
}
