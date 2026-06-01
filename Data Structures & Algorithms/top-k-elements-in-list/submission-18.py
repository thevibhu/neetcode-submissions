class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        freq = {}
        ans = []

        for x in nums:
            if x in freq:
                freq[x] = freq[x] + 1
            else:
                freq[x] = 1

        sorted_freq = dict(sorted(freq.items(), key=lambda item: item[1], reverse=True))
        val = 1
        for key, value in sorted_freq.items():
            if val <= k:
                ans.append(key)
            val = val + 1
        
        return ans
        
        