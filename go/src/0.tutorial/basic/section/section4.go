package main

import "fmt"

func main() {
	var num1, num2 int
	
	fmt.Scanln(&num1, &num2)
	num3 := (num1/num2)
	num4 := (num1%num2)
	
	fmt.Printf("몫 : %d, 나머지 : %d\n", num3, num4)
	
}