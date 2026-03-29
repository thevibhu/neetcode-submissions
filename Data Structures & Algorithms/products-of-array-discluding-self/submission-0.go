func productExceptSelf(nums []int) []int {
    n := len(nums)
    result := make([]int, n)

    // Pass 1: result[i] = product of all elements LEFT of i
    prefix := 1
    for i := 0; i < n; i++ {
        result[i] = prefix     // Store prefix BEFORE including nums[i]
        prefix *= nums[i]      // Then grow prefix
    }

    // Pass 2: multiply result[i] by product of all elements RIGHT of i
    postfix := 1
    for j := n - 1; j >= 0; j-- {
        result[j] *= postfix   // MULTIPLY, don't overwrite
        postfix *= nums[j]     // Then grow postfix
    }

    return result
}
