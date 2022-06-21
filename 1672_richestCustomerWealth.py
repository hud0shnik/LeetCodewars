#!/usr/bin/python3

class Solution:
    def maximumWealth(self, accounts: list[list[int]]) -> int:
        result=0
        for i in range(len(accounts)):
            result=max(sum(accounts[i]), result)
        return result


if __name__ == "__main__":
    s = Solution 
    print(s.maximumWealth(s,[[1,2,3],[3,2,1]]))
