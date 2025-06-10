package medium

type WaterItem struct {
	height int
	index  int
}

func maxArea(height []int) int {
	forward := make([]WaterItem, len(height))
	backward := make([]WaterItem, len(height))

	for i := 0; i < len(height); i++ {
		if i == 0 {
			forward[i] = WaterItem{
				height: height[i],
				index:  i,
			}
		} else if height[i] > forward[i-1].height {
			forward[i] = WaterItem{
				height: height[i],
				index:  i,
			}
		} else {
			forward[i] = forward[i-1]
		}
	}

	for i := len(height) - 1; i >= 0; i-- {
		if i == len(height)-1 {
			backward[i] = WaterItem{
				height: height[i],
				index:  i,
			}
		} else if height[i] > backward[i+1].height {
			backward[i] = WaterItem{
				height: height[i],
				index:  i,
			}
		} else {
			backward[i] = backward[i+1]
		}
	}

	res := 0

	for i := 0; i < len(height); i++ {
		wide := forward[i].index - backward[i].index
		if wide < 0 {
			wide = -wide
		}
		res = max(res, min(forward[i].height, backward[i].height)*wide)
	}

	return res

}
