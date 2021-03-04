package main

import "fmt"

func canIdrink(age int) bool {
	if koreanAge := age + 1; koreanAge < 18 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canIdrink(16))
}
