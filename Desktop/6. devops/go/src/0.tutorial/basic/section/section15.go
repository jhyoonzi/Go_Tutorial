package main

import "fmt"

func main() {
	
	var A = [2][2]int{
		{7,3},
		{5,2},
	}
	
	var A2 = [2][2]int{
	}
	
	d := (A[0][0] * A[1][1]) - (A[0][1] * A[1][0]) //판별식
	// 14 - 15 = -1 
	
	if d != 0 {
		fmt.Println("true")
		A2[0][0] = A[1][1]*1/d
		A2[0][1] = A[0][1]*1/d * -1
		A2[1][0] = A[1][0]*1/d * -1
		A2[1][1] = A[0][0]*1/d
		
		fmt.Println(A2[0][0],A2[0][1])
		fmt.Println(A2[1][0],A2[1][1])
		
	} else if d==0 {
		fmt.Println("false")	
	}
}