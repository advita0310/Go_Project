package features

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct{
	Name string
	PRN string
	Marks int
}

type Class struct{
	Students []Student

}

func(c *Class) NewClass(){
	myStudents := []Student{}
	file,err:=os.Open("db.txt")

	if(err!=nil){
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields:= strings.Split(line," ")
		marks,err:= strconv.Atoi(fields[2])

		if err!=nil {
			panic(err)
		}

		obj := Student{
			Name : fields[0],
			PRN : fields[1],
			Marks : marks,
		}

		myStudents=append(myStudents,obj)
	}

	fmt.Println(myStudents)

	// fmt.Println("Enter Name:")
	// var name string
	// fmt.Scanln(&name)

	// fmt.Println("Enter PRN:")
	// var PRN string
	// fmt.Scanln(&PRN)

	// fmt.Println("Enter Marks:")
	// var marks int
	// fmt.Scanln(&marks)

	// obj := Student{
	// 	Name: name,
    //     PRN: PRN,
    //     Marks: marks,
	// }

	// fmt.Println(obj)


}