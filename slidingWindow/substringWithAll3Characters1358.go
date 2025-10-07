package slidingwindow
func NumberOfSubstringsWithAllThreeChars(s string) int {
    substrings := 0 
    n := len(s)
    m := make(map[int]int)
    m[1] , m[2] , m[3] = -1 , -1 , -1
    
    for ptr := 0 ; ptr < n ; ptr++ {
        if string(s[ptr]) == "a" {
            m[1] = ptr
        } else if string(s[ptr]) == "b" {
            m[2] = ptr 
        } else {
            m[3] = ptr
        }
        if m[3] >= 0 && m[2] >= 0 && m[1] >=0 {
            substrings += min(m[1],m[2],m[3]) + 1 
        }
    } 

    return substrings

}