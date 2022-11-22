#!/usr/bin/python3

class Solution:
    def sortSentence(self, s: str) -> str:
        words = s.split()
        result = ''
        i = 1
        while i<=len(words):
            for word in words:
                if word[-1] == str(i):
                    result += word[:-1] + ' '
            i += 1
            
        return result[:-1]

if __name__ == "__main__":
    s = Solution 
    print(s.sortSentence(s,"is2 sentence4 This1 a3"))
