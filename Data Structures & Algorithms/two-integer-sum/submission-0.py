class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        hashing = {}
        for index, elem in enumerate(nums):
            if target - elem in hashing:
                return [hashing[target-elem], index]
            else:
                hashing[elem] = index