class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        
        n = len(nums)

        ans = [1] * n


        prefix = 1
        for i, val in enumerate(nums):
            ans[i] = prefix
            prefix = prefix * nums[i]        


        suffix = 1
        for i, val in reversed(list(enumerate(nums))):
            ans[i] = ans[i] * suffix
            suffix = nums[i] * suffix
        return ans
