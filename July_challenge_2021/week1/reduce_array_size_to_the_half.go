type pair struct{
    value int
    frequency int
}
type pairlist []pair

func (p pairlist) Len() int {return len(p)}
func (p pairlist) Less(i,j int) bool {return p[i].frequency<p[j].frequency}
func (p pairlist) Swap(i,j int) {p[i],p[j] = p[j], p[i]} 

func minSetSize(arr []int) int {
    frequency_list := make(map[int]int)
    for _,v := range arr{
        frequency_list[v]++
    }
    pl := make(pairlist,len(frequency_list))
    i := 0
    for k,v := range frequency_list{
        pl[i] = pair{k,v}
        i++
    }
    sort.Sort(sort.Reverse(pl))
    count := 0
    length := len(arr)
    L := length
    i =0
    for length>L/2{
        length -= pl[i].frequency
        i++
        count++
    }
    return count
}
