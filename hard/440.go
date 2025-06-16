package hard

import "math"

func findKthNumber(n int, k int) int {

	// 最高位数字 higgest
	higgest := n
	pow := 1
	for higgest >= 10 {
		higgest /= 10
		pow += 1
	}

	// 最高位之下的容量 （higgest-1） * [10^pow-1+pow]
	if k <= (higgest)*(int(math.Pow10(pow-1))-1+pow) {
		i := k / (int(math.Pow10(pow-1)) - 1 + pow)
		k = k % (int(math.Pow10(pow-1)) - 1 + pow)
		if k > higgest {
			k -= higgest
			return i*int(math.Pow10(higgest)) + k
		} else {
			return i * int(math.Pow10(k))
		}

	} else { // 最高位确认为higgest， 返回 higgest*math.Pow10(pow) + findKthNumber(n-higgest*10^pow, k-（higgest-1） * (math.Pow10(pow)-1+pow))
		return higgest*int(math.Pow10(pow)) + findKthNumber(n-higgest*(int(math.Pow10(pow))), k-(higgest)*(int(math.Pow10(pow-1))-1+pow))
	}

}
