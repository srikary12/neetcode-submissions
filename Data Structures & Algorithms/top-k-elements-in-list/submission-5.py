class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        nums.sort()
        res = {}
        for num in nums:
            if num in res:
                res[num] += 1
            else:
                res[num] = 1
        sortedDict = dict(sorted(res.items(), key=lambda item: item[1], reverse= True))
        resR = []
        for key, value in enumerate(sortedDict):
            if k <= 0:
                break
            resR.append(value)
            k-=1
        return resR
