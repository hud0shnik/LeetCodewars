#!/usr/bin/python3

class Solution:
    def balancedStringSplit(self, s: str) -> int:
        result = left = right = 0
        
        for ch in s:
            if ch == 'L':
                left += 1
            else:
                right += 1
            
            if left == right:
                left = right = 0
                result += 1
        
        return result

if __name__ == "__main__":
    s = Solution 
    print(s.balancedStringSplit(s,"RLRRLLRLRL"))
