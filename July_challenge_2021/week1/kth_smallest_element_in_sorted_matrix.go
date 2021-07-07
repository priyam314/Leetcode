// binary search used
// checks for mid, and checks if the count 
// greater than equal to less than k or not
// if k is greater than count, then our kth element 
// lies on right side of mid, else left
func kthSmallest(matrix [][]int, k int) int {
    n := len(matrix)
    matrix_min := matrix[0][0]
    matrix_max := matrix[n-1][n-1]
    for matrix_min<matrix_max{
        mid := (matrix_min + matrix_max)/2
        if k>count(mid, matrix, n){
            matrix_min = mid + 1
        }else{
            matrix_max = mid
        }
    }
    return matrix_min
}
// counts the number of elements less than or equal
// to in matrix than k(here the function argument k)
func count(k int, matrix [][]int, length int) int {
    count := 0
    for i:=0;i<length;i++{
        for j:=0;j<length;j++{
            if matrix[i][j]<=k{
                count++
            }else{
                break
            }
        }
    }
    return count
}
