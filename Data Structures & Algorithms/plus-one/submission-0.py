class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        strr = "".join(str(d) for d in digits)
        str_int = int(strr) + 1
        return list(str(str_int))