package kata

func MaximumSubarraySum(numbers []int) int {
	var localSum, maxSum int
	ln := len(numbers)
	for i := 0; i < ln; i++ {
		localSum = 0
		for j := i; j < ln; j++ {
			localSum += numbers[j]
			if localSum > maxSum {
				maxSum = localSum
			}
		}
	}
	return maxSum
}
