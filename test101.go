package main

import (
	"fmt"
	"strconv"
)

type Person interface {
	work() string
}

type Student struct {
	name string
	id   int
}

type Teacher struct {
	name string
	id   int
}

func (t Teacher) describe_teacher() string {
	return t.name + "  " + strconv.Itoa(t.id)
}

func (s Student) describe_student() string {
	return s.name + "  " + strconv.Itoa(s.id)
}

func (t Teacher) work() string {
	return strconv.Itoa(t.id) + " " + t.name
}

func (s Student) work() string {
	return strconv.Itoa(s.id) + " " + s.name
}

func describe_person(t Person) {
	switch t.(type) {
	case Teacher:
		fmt.Println(t.work())
		fmt.Println("Describe the person")
		fmt.Println(t.(Teacher).describe_teacher())
	case Student:
		fmt.Println(t.work())
		fmt.Println("Describe the person")
		fmt.Println(t.(Student).describe_student())
	default:
		fmt.Println("Default case executed")
	}
}

func main() {
	s1 := Student{"Bapan", 34}
	t1 := Teacher{"Ram", 66}
	describe_person(s1)
	describe_person(t1)
}
