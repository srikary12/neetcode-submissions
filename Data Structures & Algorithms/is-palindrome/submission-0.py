import string
class Solution:
    def isPalindrome(self, s: str) -> bool:
        s = s.translate(str.maketrans('', '', string.punctuation))
        s = s.split(" ")
        s = ("").join(s)
        s = s.lower()
        if s == s[::-1]:
            return True
        return False