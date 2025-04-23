package main

func countLargestGroup(n int) int {
	counts := make(map[int]int)
	for i := 0; i < 10; i++ {
		counts[i] = 0
	}
	count_max := 0
	count_number := 0
	for i := 1; i <= n; i++ {
		magic := calc(i)
		counts[magic] = counts[magic] + 1
		if counts[magic] == count_max {
			count_number++
		} else if counts[magic] > count_max {
			count_max = counts[magic]
			count_number = 1
		}
	}
	return count_number
}

func calc(n int) int {
	res := 0
	for n > 0 {
		res = res + (n % 10)
		n = n / 10
	}
	return res
}

func main() {

}
