class Solution:
    def isValid(self, s: str) -> bool:
        stak = []
        opp = {'{':'}','(':')','[':']'}

        if s[0] == '}' or s[0] == ']' or s[0] == ')':
            return False

        for x in s:
            if x in opp:  # Opening bracket
                stak.append(x)
            else:  # Closing bracket
                # If stack is empty or the popped opening bracket doesn't map to x
                if not stak or opp[stak.pop()] != x:
                    return False
                    
        return len(stak) == 0  