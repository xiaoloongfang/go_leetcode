package main

func smallestBeautifulString(s string, k int) string {
	bufferInt := make([]int, len(s))
	for i := range bufferInt {
		bufferInt[i] = int(s[i] - 'a')
	}
	var exit bool
	bufferInt, exit = add(bufferInt, 26, 1)
	if exit {
		return ""
	}
	flag := false
	cnt := 0
	for i := 0; i < len(bufferInt); i++ {
		if flag {
			bufferInt[i] = cnt % min(3, k)
			cnt++
			continue
		}
		if i == 0 && bufferInt[i] >= k {
			return ""
		}
		if bufferInt[i] >= k {
			bufferInt[i-1] = bufferInt[i-1] + 1
			if bufferInt[i-1] >= k {
				return ""
			}
			bufferInt[i] = 0
			flag = true
		}
	}

	for true {
		index := checkPalindromic(bufferInt)
		if index == -1 {
			return con2str(bufferInt)
		} else {
			bufferInt, exit = add(bufferInt, k, index)
			if exit {
				break
			}
		}
	}

	return ""
}

func add(bufferInt []int, k int, index int) ([]int, bool) {
	overflow := true
	for i := len(bufferInt) - index; i >= 0; i-- {
		if overflow {
			next := bufferInt[i] + 1
			if next >= k {
				overflow = true
				bufferInt[i] = 0
			} else {
				overflow = false
				bufferInt[i] = next
			}
		} else {
			break
		}
	}
	cnt := 0
	for i := index + 1; i < len(bufferInt); i++ {
		bufferInt[i] = cnt % min(3, k)
		cnt++
	}
	return bufferInt, overflow
}

func con2str(bufferInt []int) string {
	bufferChars := make([]byte, len(bufferInt))
	for i := 0; i < len(bufferInt); i++ {
		bufferChars[i] = byte(bufferInt[i] + 'a')
	}
	return string(bufferChars)
}

func checkPalindromic(s []int) int {
	if len(s) <= 1 {
		return -1
	}
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == s[i-1] {
			return i
		}
		if i >= 2 && s[i] == s[i-2] {
			return i
		}
	}
	return -1
}

func main() {
	println(smallestBeautifulString("edh", 7))

}
