package main

func find(mp map[byte]byte, ch byte) byte {
	if mp[ch] != ch {
		return find(mp, mp[ch])
	}
	return ch
}

func union(mp map[byte]byte, ch1 byte, ch2 byte) map[byte]byte {
	father_ch1 := find(mp, ch1)
	father_ch2 := find(mp, ch2)

	if father_ch1 < father_ch2 {
		mp[ch2] = father_ch1
		mp[father_ch2] = father_ch1
	} else {
		mp[ch1] = father_ch2
		mp[father_ch1] = father_ch2
	}

	return mp

}

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	mp := make(map[byte]byte)
	for ch := 'a'; ch <= 'z'; ch++ {
		mp[byte(ch)] = byte(ch)
	}

	for i := 0; i < len(s1); i++ {
		union(mp, s1[i], s2[i])
	}

	res := make([]byte, len(baseStr))
	for i := 0; i < len(baseStr); i++ {
		ch := find(mp, baseStr[i])
		res[i] = ch
	}

	return string(res)

}
