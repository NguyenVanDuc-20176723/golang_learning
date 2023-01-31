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

	// while in go
	i := 0
	for i < 5 {
		fmt.Println(i)
		if i == 3 {
			break
		}
		i++
	}
}
