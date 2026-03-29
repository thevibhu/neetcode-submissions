func twoSum(nums []int, target int) []int {
    mapper := make(map[int]int)
   
   for i := range nums {
    diff := target - nums[i]
    if val, ok := mapper[diff]; ok {
        return []int{val,i}
    } 
    mapper[nums[i]] = i
   }

   return []int{}
}
