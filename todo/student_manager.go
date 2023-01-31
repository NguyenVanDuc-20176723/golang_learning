package main

import "fmt"

type Student struct {
	name     string
	age      int
	mssv     string
	monHoc   []string
	bangDiem []map[string]interface{}
}

func fakeData() Student {
	student := Student{
		name: "duc", age: 23, mssv: "20176273",
		monHoc: []string{"Dai so", "Lap trinh C co ban"},
		bangDiem: []map[string]interface{}{
			{
				"name": "Dai so",
				"diem": 9,
			}, {
				"name": "Lap trinh C co ban",
				"diem": 10,
			},
		},
	}
	return student
}

func showStudents(students []Student) {
	total := len(students)
	if total == 0 {
		fmt.Println("Danh sach sinh vien rong.")
	} else {
		fmt.Printf("Co %d sinh vien.\n", total)

		for idx, student := range students {
			fmt.Println("Sinh vien thu", idx)
			fmt.Printf("%30s | %2d | %8s ", student.name, student.age, student.mssv)
		}
	}

}

func main() {
	var students []Student
	students = append(students, fakeData())

	fmt.Println(students)
}
