#!/usr/bin/python3

class Solution:
    def runningSum(self, nums: list[int]) -> list[int]:
        result = []
        result.append(nums[0])
        for i in range(1,len(nums)):
            result.append(result[i-1]+nums[i])
        return result


if __name__ == "__main__":
    s = Solution 
    print(s.runningSum(s,[1,1,1,1,1]))
