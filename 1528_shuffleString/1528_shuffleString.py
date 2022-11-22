#!/usr/bin/python3

class Solution:
    def restoreString(self, s: str, indices: list[int]) -> str:
        d = ['']*len(s)
        for i in range(len(s)):
            d[indices[i]] = s[i]
        return ''.join(d)

if __name__ == "__main__":
    s = Solution 
    print(s.restoreString(s,"codeleet", [4,5,6,7,0,2,1,3]))
