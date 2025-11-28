package __obf_af42fb6cde2beed6

import (
	"container/heap"
	"sort"
)

type Heap struct {
	Nodes Nodes
	K     uint32
}

func NewHeap(__obf_43d363f3f5d57e86 uint32) *Heap {
	__obf_146e082d3ec5bb18 := Nodes{}
	heap.Init(&__obf_146e082d3ec5bb18)
	return &Heap{Nodes: __obf_146e082d3ec5bb18, K: __obf_43d363f3f5d57e86}
}

func (__obf_146e082d3ec5bb18 *Heap) Add(__obf_ca7bc8bea7abb9e1 *Node) *Node {
	if __obf_146e082d3ec5bb18.K > uint32(len(__obf_146e082d3ec5bb18.Nodes)) {
		heap.Push(&__obf_146e082d3ec5bb18.Nodes, __obf_ca7bc8bea7abb9e1)
	} else if __obf_ca7bc8bea7abb9e1.Count > __obf_146e082d3ec5bb18.Nodes[0].Count {
		__obf_547b0fab4fd371a5 := heap.Pop(&__obf_146e082d3ec5bb18.Nodes)
		heap.Push(&__obf_146e082d3ec5bb18.Nodes, __obf_ca7bc8bea7abb9e1)
		__obf_3cb0586d2ab5e5f6 := __obf_547b0fab4fd371a5.(*Node)
		return __obf_3cb0586d2ab5e5f6
	}
	return nil
}

func (__obf_146e082d3ec5bb18 *Heap) Pop() *Node {
	__obf_547b0fab4fd371a5 := heap.Pop(&__obf_146e082d3ec5bb18.Nodes)
	return __obf_547b0fab4fd371a5.(*Node)
}

func (__obf_146e082d3ec5bb18 *Heap) Fix(__obf_94e510a88dfa2303 int, __obf_b4fd98c1e0169c13 uint32) {
	__obf_146e082d3ec5bb18.Nodes[__obf_94e510a88dfa2303].Count = __obf_b4fd98c1e0169c13
	heap.Fix(&__obf_146e082d3ec5bb18.Nodes, __obf_94e510a88dfa2303)
}

func (__obf_146e082d3ec5bb18 *Heap) Min() uint32 {
	if len(__obf_146e082d3ec5bb18.Nodes) == 0 {
		return 0
	}
	return __obf_146e082d3ec5bb18.Nodes[0].Count
}

func (__obf_146e082d3ec5bb18 *Heap) Find(__obf_55cde42f6d47c5be string) (int, bool) {
	for __obf_d6db326dc08be53b := range __obf_146e082d3ec5bb18.Nodes {
		if __obf_146e082d3ec5bb18.Nodes[__obf_d6db326dc08be53b].Key == __obf_55cde42f6d47c5be {
			return __obf_d6db326dc08be53b, true
		}
	}
	return 0, false
}

func (__obf_146e082d3ec5bb18 *Heap) Sorted() Nodes {
	__obf_54b8f428ecb9f031 := append([]*Node(nil), __obf_146e082d3ec5bb18.Nodes...)
	sort.Sort(sort.Reverse(Nodes(__obf_54b8f428ecb9f031)))
	return __obf_54b8f428ecb9f031
}

type Nodes []*Node

type Node struct {
	Key   string
	Count uint32
}

func (__obf_da62d4b37cd624ff Nodes) Len() int {
	return len(__obf_da62d4b37cd624ff)
}

func (__obf_da62d4b37cd624ff Nodes) Less(__obf_d6db326dc08be53b, __obf_24eb5645ee6b8c63 int) bool {
	return (__obf_da62d4b37cd624ff[__obf_d6db326dc08be53b].Count < __obf_da62d4b37cd624ff[__obf_24eb5645ee6b8c63].Count) || (__obf_da62d4b37cd624ff[__obf_d6db326dc08be53b].Count == __obf_da62d4b37cd624ff[__obf_24eb5645ee6b8c63].Count && __obf_da62d4b37cd624ff[__obf_d6db326dc08be53b].Key > __obf_da62d4b37cd624ff[__obf_24eb5645ee6b8c63].Key)
}

func (__obf_da62d4b37cd624ff Nodes) Swap(__obf_d6db326dc08be53b, __obf_24eb5645ee6b8c63 int) {
	__obf_da62d4b37cd624ff[__obf_d6db326dc08be53b], __obf_da62d4b37cd624ff[__obf_24eb5645ee6b8c63] = __obf_da62d4b37cd624ff[__obf_24eb5645ee6b8c63], __obf_da62d4b37cd624ff[__obf_d6db326dc08be53b]
}

func (__obf_da62d4b37cd624ff *Nodes) Push(__obf_ca7bc8bea7abb9e1 any) {
	*__obf_da62d4b37cd624ff = append(*__obf_da62d4b37cd624ff, __obf_ca7bc8bea7abb9e1.(*Node))
}

func (__obf_da62d4b37cd624ff *Nodes) Pop() any {
	var __obf_ca7bc8bea7abb9e1 *Node
	__obf_ca7bc8bea7abb9e1, *__obf_da62d4b37cd624ff = (*__obf_da62d4b37cd624ff)[len((*__obf_da62d4b37cd624ff))-1], (*__obf_da62d4b37cd624ff)[:len((*__obf_da62d4b37cd624ff))-1]
	return __obf_ca7bc8bea7abb9e1
}
