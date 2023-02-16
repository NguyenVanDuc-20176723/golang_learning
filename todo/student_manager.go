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

func inputStudentInfo() Student {
	student := Student{}
	fmt.Println("Nhap thong tin sinh vien:")
	fmt.Println("Nhap ten sinh vien:")
	fmt.Scanf("%s", &student.name)
	fmt.Println("Nhap tuoi sinh vien:")
	fmt.Scanf("%d", &student.age)
	fmt.Println("Nhap ma so sinh vien:")
	fmt.Scanf("%s", &student.mssv)
	return student
}

func addStudents() Student {
	student := inputStudentInfo()
	return student
}

func remove(students []Student, idx int) []Student {
	return append(students[:idx], students[idx+1:]...)
}

func deleteStudent(students []Student) []Student {
	var idx int
	fmt.Println("Nhap stt sinh vien can xoa:")
	fmt.Scanf("%d", &idx)
	return remove(students, idx)
}

func showStudents(students []Student) {
	total := len(students)
	if total == 0 {
		fmt.Println("Danh sach sinh vien rong.")
	} else {
		fmt.Printf("Co %d sinh vien.\n", total)
		fmt.Println("------------------------------------------------")
		fmt.Printf("%-3s | %-20s | %-4s | %-8s \n", "No", "Name", "Age", "MSSV")
		fmt.Println("------------------------------------------------")
		for idx, student := range students {
			fmt.Printf("%-3d | %-20s | %-4d | %-8s \n", idx, student.name, student.age, student.mssv)
		}
	}

}

func menu() int {
	var choice int
	options := []string{
		"Add students.",
		"Update student.",
		"Show list student.",
		"Delete student.",
		"Exit.",
	}
	fmt.Println("------------Menu-----------")
	for idx, option := range options {
		fmt.Printf("\t%d. %s\n", idx+1, option)
	}
	fmt.Println("----------------------------")
	fmt.Println("Your choice: ")
	fmt.Scan(&choice)
	fmt.Println("----------------------------")
	return choice
}

func main() {
	var students []Student

loop:
	for true {
		choice := menu()
		switch choice {
		case 1:
			fmt.Println("--- Add students ---")
			student := addStudents()
			students = append(students, student)
			//students = append(students, fakeData())
		case 2:
			fmt.Println("--- Update student ---")

		case 3:
			fmt.Println("--- Show list student ---")
			showStudents(students)
		case 4:
			fmt.Println("--- Delete student ---")
			students = deleteStudent(students)
		case 5:
			fmt.Println("Program exited.")
			break loop
		default:
			fmt.Println("Invalid choice.")
		}
		fmt.Println()
	}

}
