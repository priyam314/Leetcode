type position struct{
    y int
    x int
}
func maxSumSubmatrix(matrix [][]int, k int) int {
    // dp would be instantiated with row+1,col+1 size
    // all elements zero
    dp := make([][]int,len(matrix)+1)
    for i,_ := range dp{
        dp[i] = make([]int,len(matrix[0])+1)
    }   
    for i:=1;i<len(matrix)+1;i++{
        for j:=1;j<len(matrix[0])+1;j++{
            dp[i][j] = dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1] + matrix[i-1][j-1]
        }
    }
    maxKSum := -1000000
    for i:=1;i<len(dp);i++{
        for j:=1;j<len(dp[0]);j++{
            for p:=i;p<len(dp);p++{
                for l:=j;l<len(dp[0]);l++{
                    //fmt.Println("top:(",i,j,")"," bottom:(",p,l,")")
                    sum := subMatrixSum(dp,position{i,j},position{p,l})
                    //fmt.Println("sum:",sum)
                    if sum==k{
                        return k
                    }else if sum<k && sum>maxKSum{
                        maxKSum = sum
                    }
                }
            }
        }
    }
    return maxKSum
}
func subMatrixSum(dp [][]int, top position, bottom position) int {
    return dp[bottom.y][bottom.x] - dp[bottom.y][top.x-1] - dp[top.y-1][bottom.x] + dp[top.y-1][top.x-1]
}
