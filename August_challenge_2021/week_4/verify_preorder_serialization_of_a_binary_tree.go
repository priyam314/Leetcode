import "strings"
func isValidSerialization(preorder string) bool {
    node := strings.Split(preorder, ",")
    // at the starting there is one vacancy for a node
    // i.e root node
    vacancy := 1
    for _,v := range node{
        // with every iteration you lose one vacancy 
        vacancy--
        // if ## is found, then vacancy lose by two
        if vacancy<0{return false}
        // for every element, add teo vacancies
        if v!= "#"{
            vacancy += 2
        }
    }
    return vacancy==0
}
