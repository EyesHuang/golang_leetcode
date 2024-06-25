package leetcode_350

import "container/heap"

// Time complexity: O(n)
// Space complexity: O(n)
func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)

	for _, num := range nums {
		count[num]++
	}

	freq := make([][]int, len(nums)+1)
	for num, c := range count {
		freq[c] = append(freq[c], num)
	}

	res := make([]int, 0, k)
	for i := len(freq) - 1; i >= 0; i-- {
		for _, num := range freq[i] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	}

	return nil
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest, not highest, priority so we use less than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Time complexity: O(n*log(k))
// Space complexity: O(n+k)
func topKFrequent_heap(nums []int, k int) []int {
	// Special case: if k is equal to the length of nums, return nums
	if k == len(nums) {
		return nums
	}

	// 1. Build hash map: character and how often it appears
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	// 2. Initialize a priority queue (heap) to keep k top frequent elements
	pq := &PriorityQueue{}
	heap.Init(pq)

	// 3. Add elements to the heap, keeping only k top elements in the heap
	for num, freq := range count {
		heap.Push(pq, &Item{value: num, priority: freq})
		if pq.Len() > k {
			heap.Pop(pq)
		}
	}

	// 4. Build an output array
	top := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		item := heap.Pop(pq).(*Item)
		top[i] = item.value
	}

	return top
}
