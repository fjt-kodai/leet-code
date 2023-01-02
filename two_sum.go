package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 18

	output := twoSum(nums, target)
	fmt.Println(output)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, num := range nums {
		index, ok := m[target-num]
		// null, false := m[16]
		// 2, true := m[11]
		if ok {
			return []int{index, i}
		}

		m[num] = i
		// m[2] = 0
	}

	return []int{}
}
