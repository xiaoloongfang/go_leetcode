package main

import "fmt"

func smallestString(s string) string {
	charList := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		charList[i] = s[i]
	}
	if allA(s) {
		charList[len(charList)-1] = 'z'
		return string(charList)
	}
	flag := false
	for i := 0; i < len(s); i++ {
		if !flag && charList[i] > 'a' {
			flag = true
			charList[i] = charList[i] - 1
		} else if flag && charList[i] > 'a' {
			charList[i] = charList[i] - 1
		} else if flag {
			break
		}
	}
	return string(charList)
}

func allA(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] != 'a' {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(smallestString("abcd"))
	fmt.Println(smallestString("cba"))
	fmt.Println(smallestString("a"))
	fmt.Println(smallestString("aa"))
	fmt.Println(smallestString("bbb"))
	fmt.Println(smallestString("aac"))
	fmt.Println(smallestString("ab"))
}
