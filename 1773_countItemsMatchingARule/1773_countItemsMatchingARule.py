#!/usr/bin/python3

class Solution:
	def countMatches(self, items: list[list[str]], ruleKey: str, ruleValue: str) -> int: 
		result = 0
		d = {"type" : 0,"color" : 1, "name":2}

		x = d[ruleKey]

		for i in range(len(items)):
			if ruleValue == items[i][x]:
				result+=1

		return result

if __name__ == "__main__":
    s = Solution 
    print(s.countMatches(s,[["phone","blue","pixel"],["computer","silver","lenovo"],["phone","gold","iphone"]],"color","silver"))
