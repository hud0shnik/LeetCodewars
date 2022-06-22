#!/usr/bin/python3

class Solution:
    def numJewelsInStones(self, jewels: str, stones: str) -> int:
        result = 0 
        for stone in stones:
            if stone in jewels:
                result += 1
        return result

if __name__ == "__main__":
    s = Solution 
    print(s.numJewelsInStones(s,"aA","aAAbbbb"))
