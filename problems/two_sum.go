package main

import (
	"fmt"
)

func main() {
	// input
	nums := []int{2, 7, 11, 15}
	target := 9

	// output
	result := twoSum(nums, target)

	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, num := range nums {
		idx, ok := m[target-num]
		if ok {
			return []int{idx, i}
		}

		m[num] = i
	}

	return []int{}
}
