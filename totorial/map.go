package main

import "fmt"

func main() {
	var a = map[string]int{"a": 1, "b": 2}
	a["c"] = 3
	delete(a, "a")
	fmt.Println(a)

	var b = make(map[string]int)
	var c map[string]int
	b["a"] = 1
	fmt.Println(b == nil, c == nil)

	// check key of map
	d := map[string]int{"a": 1, "b": 2}
	val1, ok1 := d["a"]
	val2, ok2 := d["b"]
	val3, ok3 := d["c"]
	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)

	e := d
	e["a"] = 3
	fmt.Println(e, d)

	// map multi data type
	map_diff_type := map[string]interface{}{"a": 1, "b": "duc"}
	fmt.Println(map_diff_type)

	// nested map
	nested_map := map[string]interface{}{
		"name": "duc",
		"age":  23,
		"edu": []map[string]interface{}{
			{
				"name": "hust",
				"age":  5,
			},
			{
				"name": "dien chau 3",
				"age":  3,
			},
		},
	}

	fmt.Println(nested_map)

}
