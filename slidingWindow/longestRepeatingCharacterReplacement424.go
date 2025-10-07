package slidingwindow

func CharacterReplacement(s string, k int) int {
    m := make(map[byte]int)
    l := 0
    n := len(s)
    maxLength := 0
    maxF := 0 
    for r := 0 ; r <n  ; r++ {
        m[s[r]]++
        for _ , val := range m {
            if val > maxF {
                maxF = val
                
            }
        }
        if k < r - l + 1 - maxF {
            m[s[l]]--
            l++
        } else {
            maxLength = max(maxLength , r-l+1)
        }
    

    }
    

    return maxLength
}

func CharacterReplacementOptimal(s string, k int) int {
    arr := make([]int, 26)
    l , maxF , maxLength := 0 , 0 , 0
    for r := 0 ; r < len(s) ; r++ {
        arr[s[r] - 'A']++
        if arr[s[r] - 'A'] > maxF {
            maxF = arr[s[r] - 'A']
        }
        if (r - l + 1) - maxF > k {
            arr[s[l] - 'A']--
            l++
        }
        if r - l + 1 > maxLength {
            maxLength = r - l + 1
        }
    }
    

    return maxLength

}