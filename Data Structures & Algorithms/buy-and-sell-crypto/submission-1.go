func maxProfit(prices []int) int {
    res := 0
    minBuy := math.MaxInt32
    for _, today := range(prices){
        res = max(res, today-minBuy)
        if today < minBuy{
            minBuy = today
        }
    }
    return res
}
