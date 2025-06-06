package medium

import "sort"

func construct(str string) string {
	chars := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		chars[i] = str[i]
	}
	sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })
	return string(chars)
}

func groupAnagrams(strs []string) [][]string {

	mp := make(map[string][]string)
	for _, str := range strs {
		key := construct(str)
		val := mp[key]
		if val == nil {
			val = make([]string, 0)
			val = append(val, str)
			mp[key] = val
		} else {
			val = append(val, str)
			mp[key] = val
		}
	}

	res := make([][]string, 0)
	for _, val := range mp {
		res = append(res, val)
	}
	return res

}
