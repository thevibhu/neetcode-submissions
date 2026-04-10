func longestConsecutive(nums []int) int {

	mapper := make(map[int]int)
	longset := 0

	for i, val := range nums {
		mapper[val] = i
	}

	for _, val := range nums {
		_, ok := mapper[val-1]
		if !ok {
			l := 0
			_, k := mapper[val + l]
			for k {
				l++
				_, k = mapper[val + l]
			}
		longset = int(math.Max(float64(l),float64(longset)))
		}

	}
	return longset

}
