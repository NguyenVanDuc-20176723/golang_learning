package main

import "fmt"

func main() {
	//mySlice := []int{1, 2, 3}
	//fmt.Println(mySlice, len(mySlice), cap(mySlice))

	myArray := []int{1, 2, 3, 4}
	mySlice := myArray[1:3]
	mySlice[1] = 5
	fmt.Println(mySlice, myArray, len(mySlice), cap(mySlice))
	fmt.Println(append(myArray, 6, 7))

	mySlide1 := []int{1, 2, 3}
	mySlide2 := []int{4, 5, 6}
	fmt.Println(append(mySlide1, mySlide2...))
}
