//package main
//
//import "fmt"
//
//type Person struct {
//	Name string
//	Age  int
//}
//
//func (p Person) String() string {
//	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
//}
//
//func main() {
//	a := Person{"Arthur Dent", 42}
//	z := Person{"Zaphod Beeblebrox", 9001}
//	fmt.Println(a, z)
//}

package main

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ip IPAddr) String() string {
	var str []string
	for _, value := range ip {
		str = append(str, fmt.Sprintf("%d", value))
	}

	return fmt.Sprintf("%v", strings.Join(str, "."))
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
