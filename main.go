package main

import "container/heap"

type Node struct {
	char  rune
	freq  int
	left  *Node
	right *Node
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { return pq[i].freq < pq[j].freq }

func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x any) { *pq = append(*pq, x.(*Node)) }

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func HuffmanCoding(data string) *Node {
	freqMap := make(map[rune]int)

	for _, char := range data {
		freqMap[char]++
	}

	pq := &PriorityQueue{}
	heap.Init(pq)

	for char, freq := range freqMap {
		heap.Push(pq, &Node{char: char, freq: freq})
	}

	for pq.Len() > 1 {
		first := heap.Pop(pq).(*Node)
		second := heap.Pop(pq).(*Node)

		parent := &Node{
			freq:  first.freq + second.freq,
			left:  first,
			right: second,
		}

		heap.Push(pq, parent)
	}

	return heap.Pop(pq).(*Node)
}
