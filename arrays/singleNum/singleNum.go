package singleNum


/*
Given a non-empty array of integers nums, every element appears twice except for one. 
Find that single one.
You must implement a solution with a linear runtime complexity
 and use only constant extra space. */
func SingleNumber(nums[]int) int {
	n := make(map[int]int)
	for _, v := range nums {
		n[v] += 1
	}
	for ele, v := range n {
		if v == 1 { return ele }
	}
	return -1 
}