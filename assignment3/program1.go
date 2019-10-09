package main

import (
	"fmt"
	"strings"
)

type student struct {
	rolno int
	fname, lname string
}

type intf interface {
	display()
}

func main()  {
	s := student{
		rolno: 1,
		fname: "amol",
		lname: "pawar",
	}
	s.display()
}

func (s student)  display(){
	full_name := strings.Join([]string{s.fname,s.lname}," ")
	fmt.Printf("RollNO: %d and FullName: %s", s.rolno, full_name)
}