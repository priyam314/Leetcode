func grayCode(n int)[]int{
    output := make([]int,0)
    for v := 0;v<1<<n;v++{
        output = append(output,v^(v>>1))
    }
    return output
}
