class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        ans = {}

        for x in nums:
            if x in ans:
                ans[x] +=1
            else:
                ans[x] = 1
        
        for x, v in ans.items():
            if v ==1 :
                return x
    