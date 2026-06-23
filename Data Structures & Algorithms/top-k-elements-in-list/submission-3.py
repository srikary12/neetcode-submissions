class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        hashing = defaultdict(int)
        for elem in nums:
            hashing[elem] += 1
        resMap = sorted(hashing.items(), key = lambda item: item[1], reverse=True)
        res = []
        i=0
        while k > 0:
            res.append(resMap[i][0])
            k -= 1
            i += 1
        return res