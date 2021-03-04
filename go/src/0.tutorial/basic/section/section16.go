package main

import "fmt"

func bubbleSort(nums []int) []int {
	var temp int

	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				temp = nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
	}
	return nums
}

func inputNums() (nums []int) {
	var numScale, num int
	fmt.Println("Input")
	fmt.Scanln(&numScale)

	for i := 0; i < numScale; i++ {
		fmt.Scanln(&num)
		nums = append(nums, num)
	}
	return
}

func outputNums(nums []int) {
	fmt.Println("Output")
	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d ", nums[i])
	}
}

func main() {
	nums := inputNums()
	outputNums(bubbleSort(nums))
}
