package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	sum := twoSum(nums, 26)
	fmt.Print(sum)
}

func twoSum(nums []int, target int) []int {

	firstIndex := 0
	var lastIndex = firstIndex + 1
	for ; firstIndex < len(nums); firstIndex++ {
		for ; lastIndex < len(nums); lastIndex++ {
			if nums[firstIndex]+nums[lastIndex] == target {
				goto SUCCESS
			}
		}
	}

SUCCESS:
	var resultArray = []int{firstIndex, lastIndex}

	return resultArray
}
