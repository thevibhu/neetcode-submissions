func hasDuplicate(nums []int) bool {
    mapper := make(map[int]bool)

    for _, v := range nums {
        _ , ok := mapper[v]
        if ok {
            mapper[v] = true
            return true
        } else {
            mapper[v] = false
        }
    }
    return false
}
