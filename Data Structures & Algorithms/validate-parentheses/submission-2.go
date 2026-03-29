func isValid(str string) bool {
	s := new(stack)

	for _, val := range str {
		fmt.Println(string(s.items))
		if val == ']' || val == '}' || val == ')'  {
			if s.isEmpty() {
				return false
			}
			switch val {
				case ']':
					if s.peek() == '[' {
						s.pop()
					} else {
						return false
					}
				case ')':
				if s.peek() == '(' {
						s.pop()
					}else {
						return false
					}
				case '}':
				if s.peek() == '{' {
						s.pop()
					}else {
						return false
					}
			}
		} else {
			s.push(byte(val))
		}
	}

	return s.isEmpty()
}

type stack struct {
	items []byte
}

func (s *stack) push(val byte) {
	s.items = append(s.items, val)
}

func (s *stack) isEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

func (s *stack) peek() byte {
	return s.items[len(s.items) - 1]
}

func (s *stack) pop() {
	s.items = s.items[:len(s.items) - 1]
}