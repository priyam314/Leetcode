func matrixReshape(mat [][]int, r int, c int) [][]int {
    if len(mat)*len(mat[0])!=r*c{
        return mat
    }
    scalar := Scalar(mat)
    return reshape(scalar,r,c)
}
// [...] -> [[...],[...]]
func reshape(list []int, r int, c int) [][]int{
    newMatrix := make([][]int,r)
    for i ,_ := range newMatrix {
        newMatrix[i] = make([]int, c)
    }
    for i,_ := range list{
        newMatrix[i/c][i%c] = list[i]
    }
    return newMatrix
}
// [[...],[...]] -> [...]
func Scalar(mat[][]int)[]int{
    row, column := len(mat), len(mat[0])
    newList := make([]int,row*column)
    for i:=0;i<row;i++{
        for j:=0;j<column;j++{
            newList[column*i + j] = mat[i][j]
        }
    }
    return newList
}
