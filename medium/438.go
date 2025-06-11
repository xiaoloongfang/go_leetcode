package medium

func findAnagrams(s string, p string) []int {
	mp := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		cnt, ok := mp[p[i]]
		if !ok {
			mp[p[i]] = 1
		} else {
			mp[p[i]] = cnt + 1
		}
	}

	res := make([]int, 0)

	left, right := 0, 0

	for right < len(s) {
		cnt, ok := mp[s[right]]
		if !ok {
			right++
			left = right
		} else if cnt == 0 {
			mp[s[left]] = mp[s[left]] + 1
			left++
		} else {
			if right-left < len(p) {
				mp[s[right]] = mp[s[right]] - 1
				right++
			} else if right-left == len(p) {
				res = append(res, left)
				mp[s[left]] = mp[s[left]] + 1
				mp[s[right]] = mp[s[right]] - 1
				left++
				right++
				for right < len(s) && s[left-1] == s[right] {
					res = append(res, left)
					mp[s[left]] = mp[s[left]] + 1
					mp[s[right]] = mp[s[right]] - 1
					left++
					right++
				}
			}
		}
	}

	return res
}
