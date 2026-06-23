class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        dic = {
            "(" : ")",
            "{": "}",
            "[": "]",
        }
        for elem in s:
            if elem in "([{":
                stack.append(elem)
            elif elem in ")}]" and len(stack)>0:
                temp = stack.pop()
                if dic[temp] != elem:
                    return False
            else:
                return False
        if len(stack) == 0:
            return True
        else:
            return False