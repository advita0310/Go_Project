package main

import (
	"fmt"
	"os"
	features "student-db/Features"
)

func main() {
	fmt.Println("Hello DB")
	var myclass features.Class
	myclass.NewClass()

	for {
		fmt.Println("1.Show All Students")
        fmt.Println("2.Add Student")
        fmt.Println("3.Update Student")
        fmt.Println("4.Delete Student")
        fmt.Println("5.Exit")

        var choice int
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            myclass.ShowStudents()
        case 2:
            myclass.AddStudent()
        case 3:
            myclass.UpdateStudent()
        case 4:
            myclass.DeleteStudent()
        case 5:
            os.Exit(0)
        default:
            fmt.Println("Invalid Choice")
        }
	}

}
