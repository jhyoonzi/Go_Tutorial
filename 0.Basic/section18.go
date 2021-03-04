package main

import (
	"fmt"
	"strings"
)

func lenAndupper(name string) (length int, uppercase string) {
	defer fmt.Println("i'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLength, up := lenAndupper("jhyoonzi")
	fmt.Println(totalLength, up)
}
