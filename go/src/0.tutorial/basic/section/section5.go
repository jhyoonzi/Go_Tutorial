package main

import "fmt"

func main() {
	
	var num1, num2, num3 int
	fmt.Scanln(&num1, &num2, &num3)
	
	
	changenum1 := float32(num1)
	changenum2 := uint(num2)
	changenum3 := int64(num3)
	
	
	fmt.Printf("float32, %f\nuint, %d\nint64, %d", changenum1, changenum2, changenum3)
}