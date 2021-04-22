/*
2단부터 9단까지 출력하는 구구단입니다. 하지만 홀수 단만 출력합니다.
"%d x %d = %d\n"형태로 출력합니다.
단과 단 사이는 한 줄을 비웁니다.
n단은 n x n 까지 출력합니다. 예를 들어, 7단은 7 x 7 = 49까지 출력합니다.
*/

package main

import "fmt"

func main() {
	
	for i:=2; i<10; i++ {
		if i%2 == 0 {
		continue
		}
		for dan:=1; dan < 10; dan++ {
			if i/dan == 0 {
				continue
			}
		fmt.Printf("%d x %d = %d\n", i, dan, i*dan)	
		}
			fmt.Printf("\n")
	}
}