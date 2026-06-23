class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        visited = {}
        for index, elem in enumerate(numbers):
            if target - elem in visited:
                return [visited[target-elem], index + 1]
            visited[elem] = index + 1
        return []