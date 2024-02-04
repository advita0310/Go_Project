package features

import "fmt"

type Student struct{
	Name string
	PRN string
	Marks int
}

type Class struct{
	Students []Student

}

func(c *Class) NewClass(){
	fmt.Println("Enter Name:")
	var name string
	fmt.Scanln(&name)

	fmt.Println("Enter PRN:")
	var PRN string
	fmt.Scanln(&PRN)

	fmt.Println("Enter Marks:")
	var marks int
	fmt.Scanln(&marks)
}