class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        sd = {}
        td = {}
        for elem in s:
            if elem in sd:
                sd[elem] += 1
            else:
                sd[elem] = 1
        for elem in t:
            if elem in td:
                td[elem] += 1
            else:
                td[elem] = 1
        return sd == td