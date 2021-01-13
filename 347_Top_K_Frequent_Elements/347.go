package main

import (
	"fmt"
	"math"
)

func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	max_count := 0
	for i := 0; i < len(nums); i++ {
		counts[nums[i]]++
		max_count = int(math.Max(float64(counts[nums[i]]), float64(max_count)))
	}

	var buckets = make([][]int, max_count+1)
	for key, value := range counts {
		buckets[value] = append(buckets[value], key)
	}

	var ans = make([]int, 0)
	for i := max_count; i >= 0; i-- {
		for _, value := range buckets[i] {
			ans = append(ans, value)
			if len(ans) == k {
				return ans
			}
		}
	}
	return ans
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}
