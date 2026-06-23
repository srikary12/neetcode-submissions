class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        hashing = {}
        for elem in strs:
            elem_sort = list(elem)
            elem_sort.sort()
            elem_sort = ('').join(elem_sort)
            if elem_sort in hashing:
                hashing[elem_sort].append(elem)
            else:
                hashing[elem_sort] = [elem]
        res = []
        for key in hashing:
            res.append(hashing[key])
        return res