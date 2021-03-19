package main

import (
	"fmt"
	"leetcode_go/solution"
)

func main() {
	nums := [6]int{3, 2, 1, 5, 6, 4}
	ret := solution.FindKthLargest(nums[:], 2)
	fmt.Println(ret)
}
