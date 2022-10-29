package main

/*
https://leetcode.com/problems/find-pivot-index/
Given an array of integers nums, calculate the pivot index of this array.

The pivot index is the index where the sum of all the numbers strictly to the left of the index is equal to the sum of all the numbers strictly to the index's right.

If the index is on the left edge of the array, then the left sum is 0 because there are no elements to the left. This also applies to the right edge of the array.

Return the leftmost pivot index. If no such index exists, return -1.

Example 1:

Input: nums = [1,7,3,6,5,6]
Output: 3
Explanation:
The pivot index is 3.
Left sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11
Right sum = nums[4] + nums[5] = 5 + 6 = 11
Example 2:

Input: nums = [1,2,3]
Output: -1
Explanation:
There is no index that satisfies the conditions in the problem statement.
Example 3:

Input: nums = [2,1,-1]
Output: 0
Explanation:
The pivot index is 0.
Left sum = 0 (no elements to the left of index 0)
Right sum = nums[1] + nums[2] = 1 + -1 = 0
*/
func FindPivotIndex(nums []int) int {
	// Matias
	//Time Submitted	Status		Runtime		Memory
	//10/30/2022 08:13	Accepted	1085 ms		21 MB
	//10/30/2022 08:13	Accepted	1364 ms		39.5 MB
	//10/30/2022 08:13	Accepted	1372 ms		49.5 MB
	//10/30/2022 08:12	Accepted	918 ms		48.9 MB
	//left := strictRunningSumLeft(nums)
	//right := strictRunningSumRight(nums)
	//for i := 0; i < len(nums); i++ {
	//	if left[i] == right[i] {
	//		return i
	//	}
	//}
	//return -1

	// Internet
	//Time Submitted	Status		Runtime		Memory
	//10/30/2022 08:22	Accepted	16 ms		6.2 MB
	//10/30/2022 08:22	Accepted	37 ms		6.3 MB
	//10/30/2022 08:20	Accepted	28 ms		6.5 MB
	left := 0
	right := sum(nums)
	for i, num := range nums {
		left += nums[i]
		if left == right {
			return i
		}
		right -= num
	}
	return -1
}

func sum(nums []int) int {
	s := 0
	for _, num := range nums {
		s += num
	}
	return s
}

func strictRunningSumLeft(nums []int) []int {
	result := []int{0}
	for i := range nums[1:] {
		result = append(result, result[i]+nums[i])
	}
	return result
}

func strictRunningSumRight(nums []int) []int {
	result := []int{0}

	lastInputIndex := len(nums) - 1
	for i := range nums[1:] {
		lastOutputIndex := len(result) - 1
		result = append(
			[]int{result[lastOutputIndex-i] + nums[lastInputIndex-i]},
			result...)
	}

	return result
}

func RunningSum(nums []int) []int {

	var acc int

	for i, num := range nums {
		acc += num
		nums[i] = acc
	}
	return nums
}
