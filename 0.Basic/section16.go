/*
슬라이스 변수 names를 make 함수를 이용해 선언합니다.
사용자에게 입력받는 이름 변수 name을 string형으로 선언합니다.
이름은 엔터(개행)로 제한 없이 입력 받고 1을 입력하면 입력을 종료합니다.
같은 길이의 이름이면 가장 먼저 입력한 이름이 출력됩니다.
가장 긴 이름과 그 길이가 결괏값으로 출력됩니다.
*/

package main

import "fmt"

func main() {
	names := make([]string, 0)
	var name string

	for {
		fmt.Scan(&name)
		if name == "1" {
			break
		} else {
			names = append(names, name)
		}
	}

	var result string = names[0]

	for i := 0; i < len(names); i++ {
		if len(result) < len(names[i]) {
			result = names[i]
		}
	}
	fmt.Println(result, len(result))
}
