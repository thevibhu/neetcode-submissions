func search(nums []int, target int) int {
    l := 0
    h := len(nums) -1

    return searchRec(nums, target, l, h)
}

func searchRec(nums []int, target int, l int, h int) int {
    if l > h {
        return -1
    }

    mid := int((l + h)/2)
    if nums[mid] > target {
        return searchRec(nums, target, l, mid-1)
    } else if nums[mid] < target{
        return searchRec(nums, target, mid + 1, h)
    } else{
        return mid
    }
}