func findLength(nums1 []int, nums2 []int) int {
    // initializing matrix of len(nums1)+1,len(nums2)+1
    dp := make([][]int, len(nums1)+1)
    for i,_ := range dp{
        dp[i] = make([]int, len(nums2)+1)
    }
    row := len(nums1)
    col := len(nums2)
    maxCount := 0
    
    // adding 1 to dp[i-1][j-1] 
    // where nums1 and nums2 element equal
    // keeping track of largest dp element
    // most of the elements are zero in dp
    // so lots of unessential iteration is happening
    for i:=1;i<row+1;i++{
        for j:=1;j<col+1;j++{
            if nums1[i-1]==nums2[j-1]{
                dp[i][j] = 1 + dp[i-1][j-1]
                if dp[i][j]>maxCount{
                    maxCount = dp[i][j]
                }
            }else{
                dp[i][j] = 0
            }
        }
    }
    return maxCount
}
