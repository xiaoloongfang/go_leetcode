package hard

type CharItem struct {
	char   byte
	needed int
	indexs []int
}

func (this *CharItem) append(index int) {
	this.indexs = append(this.indexs, index)
}

func (this *CharItem) pop(index int) {
	for i := 0; i < len(this.indexs); i++ {
		if this.indexs[i] > index {
			this.indexs = this.indexs[i : len(this.indexs)-1]
		}
	}
	this.indexs = []int{}
}

func check(mp map[byte]CharItem) bool {
	for _, charItem := range mp {
		if len(charItem.indexs) < charItem.needed {
			return false
		}
	}
	return true
}

func minWindow(s string, t string) string {
	mp := make(map[byte]CharItem)
	for index, ch := range t {
		charItem, ok := mp[byte(ch)]
		if !ok {
			charItem = CharItem{char: byte(ch), needed: 1, indexs: []int{index}}
			mp[byte(ch)] = charItem
		} else {
			// charItem.indexs = append(charItem.indexs, index)
			charItem.needed++
		}
	}

	startIndex := 0
	res := ""

	for i := 0; i < len(s); i++ {
		ch := byte(s[i])
		charItem := mp[ch]
		charItem.append(i)
		mp[ch] = charItem

		if check(mp) {

		} else {

		}
	}

	return res

}
