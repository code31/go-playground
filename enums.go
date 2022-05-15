package main

import "fmt"

// From npf.io safer enums blog post

type MyEnum struct {
	name string
}

func (e MyEnum) String() string { return e.name }

var (
	Test1 = MyEnum{"Test1"}
	Test2 = MyEnum{"Test2"}
)

func IsValid(e MyEnum) string {
	return fmt.Sprintf("%s", e)
}

func main() {
	valid1 := IsValid(Test1)
	valid2 := IsValid(Test2)
	//notValid := IsValid("not valid")

	fmt.Println(valid1, valid2)
}
