package main

import "fmt"

func xs() {
	xs := []float64{98, 93, 77, 82, 83}

	total := 0.0
	for _, v := range xs {
		total += v
	}
	fmt.Println(total / float64(len(xs)))
}

func average(xs []float64) float64 {
	panic("Not Implenteed")
}
