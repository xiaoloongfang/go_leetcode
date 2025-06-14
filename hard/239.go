package hard

type HeapItem struct {
	index int
	value int
}

type Heap struct {
	items []HeapItem
	cap   int
	size  int
}

func (h *Heap) Pop() HeapItem {
	h.size--
	h.items[0], h.items[h.size] = h.items[h.size], h.items[0]
	index := 0
	left, right := 2*index+1, 2*index+2
	for left < h.size {
		if right < h.size && h.items[right].value > h.items[left].value {
			h.items[index], h.items[right] = h.items[right], h.items[index]
			index = right
			left, right = 2*index+1, 2*index+2
		} else {
			h.items[index], h.items[left] = h.items[left], h.items[index]
			index = left
			left, right = 2*index+1, 2*index+2
		}
	}
	return h.items[h.size]
}

func (h *Heap) Push(item HeapItem) {
	h.items[h.size] = item
	index := h.size
	h.size++
	father := (index - 1) / 2
	for index > 0 && h.items[father].value < h.items[index].value {
		h.items[index], h.items[father] = h.items[father], h.items[index]
		index = father
		father = (index - 1) / 2
	}
}

func (h *Heap) Peek() HeapItem {
	return h.items[0]
}

func (h *Heap) Init(cap int) {
	h.items = make([]HeapItem, cap)
	h.size = 0
	h.cap = cap
}

func maxSlidingWindow(nums []int, k int) []int {
	var h Heap
	h.Init(len(nums) + 1)
	for i := 0; i < k; i++ {
		h.Push(HeapItem{index: i, value: nums[i]})
	}
	res := make([]int, len(nums)-k+1)
	res[0] = h.Peek().value
	for i := k; i < len(nums); i++ {
		h.Push(HeapItem{index: i, value: nums[i]})
		for h.Peek().index <= i-k {
			h.Pop()
		}
		res[i-k+1] = h.Peek().value
	}
	return res
}
