package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	//variable
	var a int = 1
	varx := 1
	cong := a + varx
	nhan := a * varx
	if false {
		fmt.Println("Ket qua lon hon hoac bang 2")
	} else {
		fmt.Println("Ket qua nho hon 2")
	}
	fmt.Println("cong 2 so", cong, "nhan 2 so", nhan)
}
