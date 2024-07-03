package main

func removeTrailingZeros(num string) string {
	index := len(num) - 1
	for ; index >= 0 && num[index] == '0'; index-- {
	}
	return num[:index]
}
