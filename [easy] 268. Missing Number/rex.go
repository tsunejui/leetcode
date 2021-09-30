package main

func missingNumber(nums []int) int {
	var sum int
	length := len(nums)
	total := length
	for i := 0; i < length; i++ {
		total += i
		sum += nums[i]
	}
	return total - sum
}
