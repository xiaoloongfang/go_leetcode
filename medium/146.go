package medium

type CacheUnit struct {
	prv   *CacheUnit
	next  *CacheUnit
	key   int
	value int
}

type LRUCache struct {
	capacity int
	size     int
	head     *CacheUnit
	last     *CacheUnit
	mp       map[int]*CacheUnit
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		size:     0,
		head: &CacheUnit{
			prv:  nil,
			next: nil,
		},
		last: nil,
		mp:   make(map[int]*CacheUnit),
	}

}

func (this *LRUCache) Get(key int) int {
	cacheUnit, ok := this.mp[key]
	if !ok {
		return -1
	}

	if this.head.next == cacheUnit {
		return cacheUnit.value
	}

	if this.last == cacheUnit {
		this.last = cacheUnit.prv
	} else {
		cacheUnit.next.prv = cacheUnit.prv
		cacheUnit.prv.next = cacheUnit.next
	}

	cacheUnit.prv = this.head
	cacheUnit.next = this.head.next
	this.head.next.prv = cacheUnit
	this.head.next = cacheUnit

	return cacheUnit.value

}

func (this *LRUCache) Put(key int, value int) {
	preValue := this.Get(key)
	if preValue >= 0 {
		this.head.next.value = value
	} else {
		cacheUnit := &CacheUnit{
			key:   key,
			value: value,
			prv:   nil,
			next:  nil,
		}
		this.mp[key] = cacheUnit
		if this.size == 0 {
			this.head.next = cacheUnit
			cacheUnit.prv = this.head
			this.last = cacheUnit
			this.size++
		} else if this.size < this.capacity {
			cacheUnit.next = this.head.next
			this.head.next.prv = cacheUnit
			this.head.next = cacheUnit
			cacheUnit.prv = this.head
			this.size++
		} else {
			cacheUnit.next = this.head.next
			this.head.next.prv = cacheUnit
			this.head.next = cacheUnit
			cacheUnit.prv = this.head

			delete(this.mp, this.last.key)
			this.last = this.last.prv

		}
	}

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
