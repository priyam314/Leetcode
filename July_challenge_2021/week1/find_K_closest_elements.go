func findClosestElements(arr []int, k int, x int) []int {
    i, j := 0, len(arr)-k
    mid := 0
    for i<j {
        mid = (i+j)/2
        if x-arr[mid]>arr[mid+k]-x{
            i = mid +1
        }else{
            j = mid
        }
    }
    return arr[i:i+k]
}
