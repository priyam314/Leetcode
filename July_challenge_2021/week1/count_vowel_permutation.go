type container struct{
    n int
    character rune
}
var dp = make(map[container]int)
func solve(n int, character rune) int{
    if value, ok := dp[container{n,character}];ok{
        return value
    }
    if n==1{
        return 1
    }
    var result int
    switch character{
    case 'a':
        result += solve(n-1,'e') % 1000000007
    case 'e':
        result += solve(n-1,'a') + solve(n-1,'i') % 1000000007
    case 'i':
        result += solve(n-1,'a') + solve(n-1,'e') + solve(n-1,'o') + solve(n-1,'u') % 1000000007
    case 'o':
        result += solve(n-1,'i') + solve(n-1,'u') % 1000000007
    case 'u':
        result += solve(n-1,'a') % 1000000007   
    }
    dp[container{n,character}] = result
    return result
}
func countVowelPermutation(n int) int {
    a := solve(n,'a')
    e := solve(n,'e')
    i := solve(n,'i')
    o := solve(n,'o')
    u := solve(n,'u')
    return (a+e+i+o+u)%1000000007
}
