package main

import "fmt"

func main() {
	var result int
	
	for a:=0; a<10; a++{
		for b:=0; b<10; b++{
			
			result = (a*10 + b) + (b*10 + a)
			if result != 99 {
				continue
			}
			fmt.Printf("%d%d + %d%d = %d\n",a,b,b,a,result)
		}
	}
}