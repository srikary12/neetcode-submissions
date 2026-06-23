class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        hashing = []
        for elem in nums:
            if elem in hashing:
                return True
            else:
                hashing.append(elem)
        return False