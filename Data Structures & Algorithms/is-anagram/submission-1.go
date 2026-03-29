func isAnagram(s string, t string) bool {
     if len(s) != len(t) {
        return false
    }
    
    sMap := make(map[byte]int)
    tMap := make(map[byte]int)

    for  i :=0;i<len(s); i++ {
        val, ok := sMap[s[i]]
        if ok {
            sMap[s[i]]= val + 1
        } else {
            sMap[s[i]]=1
        }

        vall, ok := tMap[t[i]]
        if ok {
            tMap[t[i]]= vall + 1
        } else {
            tMap[t[i]]=1
        }
    }

     for  i :=0;i<len(t); i++ {
       val1, _ := sMap[s[i]]
       val2, _ := tMap[s[i]]
       if val1 != val2{
        return false
       } 
    }

    return true
}
