package main

func robotWithString(s string) string {
	smallest := make([]byte, len(s))
	robotString := make([]byte, len(s))
	robotIndex := -1
	res := make([]byte, 0)

	for i := len(s) - 1; i >= 0; i-- {
		if i == len(s)-1 {
			smallest[i] = s[i]
			continue
		}
		robotString[i] = min(s[i], smallest[i+1])
	}

	for i := 0; i < len(s); i++ {

		for robotIndex >= 0 && robotString[robotIndex] <= smallest[i] {
			res = append(res, robotString[robotIndex])
			robotIndex--
		}
		if s[i] > smallest[i] {
			robotIndex++
			robotString[robotIndex] = s[i]
		} else {
			res = append(res, s[i])
		}
	}

	for robotIndex >= 0 {
		res = append(res, robotString[robotIndex])
		robotIndex--
	}
	return string(res)

}
