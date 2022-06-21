#!/usr/bin/python3

class Solution:
    def buildArray(self, nums: list[int]) -> list[int]:
        answer = []
        for i in range(len(nums)):
            answer.append(nums[nums[i]])
        return answer


if __name__ == "__main__":
    s = Solution 
    print(s.buildArray(s,[0,2,1,5,3,4]))
