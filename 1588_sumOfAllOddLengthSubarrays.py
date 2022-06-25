#!/usr/bin/python3

class Solution:
    def sumOddLengthSubarrays(self, arr: list[int]) -> int:
        result = list()
        size = len(arr)
        step = 1
        jump = size if size%2 else size-1

        while step<=jump:
            for i in range(size):
                if i+step>=size+1:
                    continue
                else:
                    result.append(sum(arr[i:i+step]))
            step +=2

        return sum(result)

if __name__ == "__main__":
    s = Solution 
    print(s.sumOddLengthSubarrays(s,[1,4,2,5,3]))
