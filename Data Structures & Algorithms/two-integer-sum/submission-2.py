class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        vis = {}
        for i,v in enumerate(nums):
            if target-v in vis:
                return [vis[target-v], i]
            else:
                vis[v] = i
        return None