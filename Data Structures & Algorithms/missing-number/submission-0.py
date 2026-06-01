class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        ans ={}

        for x in nums:
            if x in ans:
                ans[x] +=1
            else:
                ans[x] = 1
        
        for i in range(1000000000000000000000):
            if i not in ans:
                return i
        