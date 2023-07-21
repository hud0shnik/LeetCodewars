package main

func createTargetArray(nums []int, index []int) []int {

	res := make([]int, len(nums))

	for k, i := range index {

		for j := len(res) - 1; j > i; j-- {
			res[j] = res[j-1]
		}

		res[i] = nums[k]
	}

	return res
}
