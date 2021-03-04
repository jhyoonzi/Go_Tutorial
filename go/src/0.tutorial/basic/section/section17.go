package main

import (
	"fmt"
	"strings"
)

func lenAndupper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	totalLength, upperName := lenAndupper("jhyoonzi")
	fmt.Println(totalLength, upperName)
}
