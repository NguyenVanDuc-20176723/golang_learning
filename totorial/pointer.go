package main

import "fmt"

func main() {
	var p *int
	fmt.Println(p)
	i := 23
	p = &i
	fmt.Println(*p, p, &p)
	*p = 21
	fmt.Println(*p, p, &p)

	var p2 **int
	fmt.Println(p)
	a := 12
	var b *int
	b = &a
	p2 = &b
	fmt.Println(**p2, *p2, p2, &p2)

}
