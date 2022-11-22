#!/usr/bin/python3

class Solution:
    def mostWordsFound(self, sentences: list[str]) -> int:
        result = 0
        for string in sentences:
            words = string.split()
            result = max(len(words), result)

        return result


if __name__ == "__main__":
    s = Solution 
    print(s.mostWordsFound(s,["Величайшие преступления совершаются из за стремления к избытку, а не к предметам первой необходимости.","ad astra per aspera"]))
