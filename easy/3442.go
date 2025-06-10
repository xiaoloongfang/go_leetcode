package easy

func maxDifference(s string) int {
	mp := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		cnt, ok := mp[s[i]]
		if !ok {
			mp[s[i]] = 1
		}
		mp[s[i]] = cnt + 1
	}

	odd_max := -1
	odd_min := len(s)
	even_max := -1
	even_min := len(s)
	for _, cnt := range mp {
		if cnt%2 == 0 {
			even_max = max(even_max, cnt)
			even_min = min(even_min, cnt)
		} else {
			odd_max = max(odd_max, cnt)
			odd_min = min(odd_min, cnt)
		}
	}

	return max(even_max-odd_min, odd_max-even_min)
}
