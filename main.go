package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(arr))
	test()
}

func rob(nums []int) int {
	last := len(nums) - 1

	for i := last - 1; i >= 0; i-- {

		pre := i + 2
		if pre < len(nums)-1 {
			nums[i] = nums[i] + nums[pre]
		}

		if nums[i+1] > nums[i] {
			nums[i] = nums[i+1]
		}
	}
	return nums[0]
}

func test() string {
	return ""
}

func mostPoints(questions [][]int) int64 {
	a := len(questions) - 1
	maps := make([]int, len(questions))
	maps[a] = questions[a][0]

	for i := a - 1; i >= 0; i-- {

		maps[i] = questions[i][0]

		if next := i + questions[i][1] + 1; next <= a {
			maps[i] += maps[next]
		}

		if maps[i+1] > maps[i] {
			maps[i] = maps[i+1]
		}

	}

	return int64(maps[0])
}
