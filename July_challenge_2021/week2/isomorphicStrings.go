
func isIsomorphic(s string, t string) bool {
    lengths := len(s)
    alphabets := make(map[byte]byte)
    boolean := make(map[byte]bool)
    for i:=0;i<lengths;i++{
        if boolean[t[i]]==false{
            if alphabets[s[i]]==0{
                alphabets[s[i]] = t[i]
            }else{  return false    }
            boolean[t[i]] = true 
        }else if alphabets[s[i]]!=t[i]{  return false   }
    }
    return true
}
