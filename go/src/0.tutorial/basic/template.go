package main

import "fmt"

func temp() {
	var actors [4]string = [4]string{"조수현", "박건령", "이지훈", "황선우"}

	for _, actor := range actors {
		fmt.Printf("제가 좋아하는 배우는 %s 입니다.\n", actor)
	}
}
