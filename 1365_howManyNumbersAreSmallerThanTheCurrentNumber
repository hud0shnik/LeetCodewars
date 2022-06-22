#!/usr/bin/python3

class Solution:
    def smallerNumbersThanCurrent(self, nums: list[int]) -> list[int]:
        sortedNums = sorted(nums)
        for i,num in enumerate(nums):
            nums[i] = sortedNums.index(num)
        return nums

if __name__ == "__main__":
    s = Solution 
    print(s.smallerNumbersThanCurrent(s,[8,1,2,2,3]))
