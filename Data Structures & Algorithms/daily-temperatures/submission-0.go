func dailyTemperatures(temperatures []int) []int {
	st := new(stack)
	result := make([]int, len(temperatures))

	for i, currentTemp := range temperatures {
		// 1. While stack is NOT empty AND current day is warmer
		// than the day at the top of the stack
		for !st.isEmpty() && currentTemp > temperatures[st.peek()] {
			
			// 2. We found a warmer day for the index at the top of the stack
			pastIndex := st.peek()
			st.pop()
			
			// 3. The difference in days is the answer
			result[pastIndex] = i - pastIndex
		}
		
		// 4. Push the current INDEX onto the stack (not the temperature!)
		st.push(i)
	}

	return result
}

type stack struct {
	items []int
}

func (s *stack) push(val int) {
	s.items = append(s.items, val)
}

func (s *stack) pop() {
	s.items = s.items[:len(s.items)-1]
}

func (s *stack) peek() int {
	return s.items[len(s.items)-1]
}

// Don't forget to add this back so you don't panic!
func (s *stack) isEmpty() bool {
	return len(s.items) == 0
}
