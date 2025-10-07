package slidingwindow


func LengthOfLongestSubstring(s string) int {
    hashMap := make(map[byte]int)

    l := 0 
    longest := 0
    for r := 0 ; r < len(s) ; r++ {
        ch := s[r]
        if prevIndex, found := hashMap[ch]; found && prevIndex >= l{
            l = prevIndex + 1
        }
        hashMap[ch] = r

        longest = max(longest, r - l + 1)
    }
    return longest
}