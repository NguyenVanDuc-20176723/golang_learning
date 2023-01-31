package main

import "fmt"

func main() {
	//for i := 0; i < 5; i++ {
	//	fmt.Println(i)
	//}

	fruits := []string{"apple", "orange", "banana"}
	for idx, val := range fruits {
		fmt.Println(idx, val)
	}
}
