class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        temp = set()
        for elem in nums:
            if elem in temp:
                return True
            else:
                temp.add(elem)
        return False