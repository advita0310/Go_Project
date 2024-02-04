package features

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
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
	c.Students=myStudents
	fmt.Println(myStudents)

	

}

func (c *Class) ShowStudents(){
	for _,val := range c.Students{
		name := val.Name
		prn := val.PRN
		marks := val.Marks
		fmt.Println(name," ",prn," ",marks)
	}
}

func (c *Class) AddStudent(){
	var wg sync.WaitGroup
	fmt.Println("Enter Name:")
	var name string
	fmt.Scanln(&name)

	fmt.Println("Enter PRN:")
	var prn string
	fmt.Scanln(&prn)

	fmt.Println("Enter Marks:")
	var marks int
	fmt.Scanln(&marks)

	obj := Student{
		Name: name,
        PRN: prn,
        Marks: marks,
	}

	c.Students = append(c.Students,obj)

	//Writing data to the file
	wg.Add(1)
	go c.WriteToFile(&wg)
	fmt.Println("Data Added")

}

func (c *Class) UpdateStudent(){
    fmt.Println("Enter PRN:")
    var prn string
    fmt.Scanln(&prn)

    for i,val := range c.Students{
        if val.PRN == prn{
            fmt.Println("Enter Name:")
			var name string
			fmt.Scanln(&name)

			fmt.Println("Enter Marks:")
			var marks int
			fmt.Scanln(&marks)

			fmt.Println("Enter PRN:")
			var PRN string
			fmt.Scanln(&PRN)
			c.Students[i]=Student{Name: name,PRN: PRN,Marks: marks}
        }
    }
	fmt.Println("Data Updated")
}

func (c *Class) DeleteStudent(){
	fmt.Println("Enter PRN:")
    var prn string
    fmt.Scanln(&prn)
	var idx int
    for i,val := range c.Students{
        if val.PRN == prn{
           idx = i;
		   break; 
        }
    }

	tempArr:=c.Students

	tempArr=append(tempArr[:idx],tempArr[idx+1:]...)

	c.Students=tempArr

	fmt.Println("Data Deleted")
}

func (c *Class) WriteToFile(wg *sync.WaitGroup){
	defer wg.Done()

	var data []string

	for _,v:=range c.Students{
		line:=fmt.Sprintf("%v %v %v",v.Name,v.PRN,v.Marks)
		//fmt.Println(line)
		data=append(data,line)
	}

	err:=os.WriteFile("db.txt",[]byte(strings.Join(data,"\n")),0664)
	if err!=nil{
        fmt.Println(err.Error())
    }
	fmt.Println("Data Written")
}