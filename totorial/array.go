package main

import "fmt"

func main() {
	var arr1 = [4]int{1, 2, 3, 4}
	fmt.Println(arr1)

	var arr2 = [...]string{"a", "b", "c"}
	fmt.Println(arr2)

	arr3 := [...]float64{1.0, 2.0, 3.0, 4.0}
	fmt.Println(arr3)

	arr4 := [...]int{1: 10, 3: 20}
	fmt.Println(arr4, len(arr4))

	// multi data type
	arr5 := [...]interface{}{1, "duc"}
	fmt.Println(arr5)
}
