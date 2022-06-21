#!/usr/bin/python3

class Solution:
    def finalValueAfterOperations(self, operations: list[str]) -> int:
        result = 0
        for command in operations:
            if '-' in command:
                result -= 1
            else:
                result += 1
                
        return result


if __name__ == "__main__":
    s = Solution 
    print(s.finalValueAfterOperations(s,["--X","X++","X--","X++","++X"]))
