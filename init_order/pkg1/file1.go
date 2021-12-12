package pkg1

import (
	"fmt"

	_ "example/pkg2"
)

var (
	_        = constInitCheck_file1()
	file1_v1 = variableInit_file1("file1_v1")
	file1_v2 = variableInit_file1("file1_v2")
)

const (
	file1_c1 = "c1"
	file1_c2 = "c2"
)

func constInitCheck_file1() string {
	if file1_c1 != "" {
		fmt.Println("pkg1: const file1_c1 has been initialized")
	}
	if file1_c2 != "" {
		fmt.Println("pkg1: const file1_c2 has been initialized")
	}
	return ""
}
func variableInit_file1(name string) string {
	fmt.Printf("pkg1: var %s has been initialized\n", name)
	return name
}
func init() {
	fmt.Println("pkg1: file1 first init func invoked")
}
func init() {
	fmt.Println("pkg1: file1 second init func invoked")
}
