#!/usr/bin/python3

class Solution:
    def defangIPaddr(self, address: str) -> str:
        address = address.replace(".","[.]")
        return address


if __name__ == "__main__":
    s = Solution 
    print(s.defangIPaddr(s,"127.0.0.1"))
