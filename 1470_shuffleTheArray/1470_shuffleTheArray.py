#!/usr/bin/python3

class Solution:
    def shuffle(self, nums: list[int], n: int) -> list[int]:
        result=[]
        for i in range(len(nums)):
            if i+n<=len(nums)-1:
                result.append(nums[i])
                result.append(nums[i+n])
        return result

if __name__ == "__main__":
    s = Solution 
    print(s.shuffle(s,[2,5,1,3,4,7],3))
