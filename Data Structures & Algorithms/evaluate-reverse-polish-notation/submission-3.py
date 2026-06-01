class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        stak = []
        operators = {"+", "-", "*", "/"}


        for x in tokens:
            if x not in operators:
                stak.append(int(x))
            else:
                b = int(stak.pop())
                a = int(stak.pop())
                if x == "+":
                    stak.append(a + b)
                elif x == "-":
                    stak.append(a - b)
                elif x == "*":
                    stak.append(a * b)
                elif x == "/":
                    stak.append(int(a / b))
        return stak[0]
        