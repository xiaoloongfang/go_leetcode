package medium

func lengthOfLongestSubstring(s string) int {

	left, right := 0, 0
	mp := make(map[byte]bool)
	count := 0
	res := 0

	for right < len(s) {
		exists, ok := mp[s[right]]
		if !ok || !exists {
			mp[s[right]] = true
			count++
			res = max(res, count)
			right++
		} else {
			mp[s[left]] = false
			left++
			count--
		}
	}

	return res

}
