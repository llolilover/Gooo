package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	result := make([]int, 0, 2)
	m := make(map[int]int)
	for i, v := range nums {
		m[v] = i

	}
	for i, v := range nums {
		value := target - v
		if i2, ok := m[value]; ok && i != i2 {
			result = append(result, i, i2)
			return result

		}
	}
	return nil

}
