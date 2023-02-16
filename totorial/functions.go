package main

import "fmt"

func show() {
	fmt.Println("Lap trinh go")
}

func myName(name string) string {
	fmt.Println(name)
	return name
}

func multiName(name1 string, name2 string) (string, string) {
	return name1, name2
}

func sum(num int) int {
	if num == 0 {
		return 0
	}
	return num + sum(num-1)
}

func validateName(name string) string {
	return name
}

//func main() {
//	show()
//	name := myName("duc")
//	fmt.Println(name)
//	a, b := multiName("Duc", "nguyen")
//	fmt.Println(a, b)
//
//	fmt.Println(sum(5))
//
//}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
