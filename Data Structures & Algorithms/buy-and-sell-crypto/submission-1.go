func maxProfit(prices []int) int {
    l, r := 0, 1
    maxP := 0.0

    for r < len(prices){
        if prices[l] < prices[r] {
            profit := prices[r] - prices[l]
            maxP = math.Max(maxP, float64(profit))
        } else {
            l =r
        }
        r++
    }

    return int(maxP)
}
