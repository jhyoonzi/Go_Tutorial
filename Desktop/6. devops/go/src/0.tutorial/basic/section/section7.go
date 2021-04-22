package main

import "fmt"

func main() {
	var RRNf,RRNt int
	var name string
	var height float32
	
	fmt.Scanf("%d-%d",&RRNf,&RRNt)
	fmt.Scan(&name)
	fmt.Scan(&height)
	
	fmt.Printf("주민등록번호 앞자리는 %d, 뒷자리는 %07d, 이름은 %s입니다.\n",RRNf, RRNt, name)
	fmt.Printf("그리고 키는 %.2f입니다.",height)
	
}