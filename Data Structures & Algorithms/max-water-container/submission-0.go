func maxArea(heights []int) int {
	l := 0
	r := len(heights) - 1
	maxArea := (r - l) * min(heights[l], heights[r])
	newArea := 0

	for l < r {
		newArea = (r - l) * min(heights[l], heights[r])
		if newArea >= maxArea{
		maxArea = newArea
		}
		if heights[l] < heights[r]{
			l++
		} else {
			r--
		}
	}

	

	return maxArea
}



func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}