func evalRPN(tokens []string) int {

 s := new(stack)

 for _, v := range tokens {
	 if isOperand(v) {
		res := s.operate(v)
		s.push(res)
	 } else {

		val, _ := strconv.Atoi(v)
		s.push(int(val))

	 }

 }
	return s.items[0]

}


type stack struct {
	items []int
}

func (s *stack) push(val int) {
	s.items = append(s.items, val)
}

func (s *stack) pop() int {
	last := s.items[len(s.items) - 1]
	s.items = s.items[:len(s.items) - 1]
	return last
}

func isOperand(b string) bool {
	if b == "+" || b == "-"  || b == "*" || b == "/" {
		return true
	}
	return false
}

func (s *stack) operate(operand string) int {
	op1 := s.pop()
	op2 := s.pop()

	switch operand {
		case "*":
			return op2 * op1
		case "-":
			return op2 - op1
		case "+":
			return op2 + op1
		case "/":
			return op2 / op1
	}

	return 0
}