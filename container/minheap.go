package __obf_e72ce603d10d02a1

import (
	"container/heap"
	"sort"
)

type Heap struct {
	Nodes Nodes
	K     uint32
}

func NewHeap(__obf_cce6a084cb2eec1f uint32) *Heap {
	__obf_41a6f945d68bc803 := Nodes{}
	heap.Init(&__obf_41a6f945d68bc803)
	return &Heap{Nodes: __obf_41a6f945d68bc803, K: __obf_cce6a084cb2eec1f}
}

func (__obf_41a6f945d68bc803 *Heap) Add(__obf_e209a94366a528e6 *Node) *Node {
	if __obf_41a6f945d68bc803.K > uint32(len(__obf_41a6f945d68bc803.Nodes)) {
		heap.Push(&__obf_41a6f945d68bc803.Nodes, __obf_e209a94366a528e6)
	} else if __obf_e209a94366a528e6.Count > __obf_41a6f945d68bc803.Nodes[0].Count {
		__obf_2795fcedd1bfa8ed := heap.Pop(&__obf_41a6f945d68bc803.Nodes)
		heap.Push(&__obf_41a6f945d68bc803.Nodes, __obf_e209a94366a528e6)
		__obf_9a0b6bec38443261 := __obf_2795fcedd1bfa8ed.(*Node)
		return __obf_9a0b6bec38443261
	}
	return nil
}

func (__obf_41a6f945d68bc803 *Heap) Pop() *Node {
	__obf_2795fcedd1bfa8ed := heap.Pop(&__obf_41a6f945d68bc803.Nodes)
	return __obf_2795fcedd1bfa8ed.(*Node)
}

func (__obf_41a6f945d68bc803 *Heap) Fix(__obf_4aa7046a273a4333 int, __obf_34ac5ddd5470ea6e uint32) {
	__obf_41a6f945d68bc803.
		Nodes[__obf_4aa7046a273a4333].Count = __obf_34ac5ddd5470ea6e
	heap.Fix(&__obf_41a6f945d68bc803.Nodes, __obf_4aa7046a273a4333)
}

func (__obf_41a6f945d68bc803 *Heap) Min() uint32 {
	if len(__obf_41a6f945d68bc803.Nodes) == 0 {
		return 0
	}
	return __obf_41a6f945d68bc803.Nodes[0].Count
}

func (__obf_41a6f945d68bc803 *Heap) Find(__obf_9dd08479fe906662 string) (int, bool) {
	for __obf_c03cf6504cc3cf88 := range __obf_41a6f945d68bc803.Nodes {
		if __obf_41a6f945d68bc803.Nodes[__obf_c03cf6504cc3cf88].Key == __obf_9dd08479fe906662 {
			return __obf_c03cf6504cc3cf88, true
		}
	}
	return 0, false
}

func (__obf_41a6f945d68bc803 *Heap) Sorted() Nodes {
	__obf_74394b9a094b4148 := append([]*Node(nil), __obf_41a6f945d68bc803.Nodes...)
	sort.Sort(sort.Reverse(Nodes(__obf_74394b9a094b4148)))
	return __obf_74394b9a094b4148
}

type Nodes []*Node

type Node struct {
	Key   string
	Count uint32
}

func (__obf_59eea0bd834c59bb Nodes) Len() int {
	return len(__obf_59eea0bd834c59bb)
}

func (__obf_59eea0bd834c59bb Nodes) Less(__obf_c03cf6504cc3cf88, __obf_c9b41611d2cb9abb int) bool {
	return (__obf_59eea0bd834c59bb[__obf_c03cf6504cc3cf88].Count < __obf_59eea0bd834c59bb[__obf_c9b41611d2cb9abb].Count) || (__obf_59eea0bd834c59bb[__obf_c03cf6504cc3cf88].Count == __obf_59eea0bd834c59bb[__obf_c9b41611d2cb9abb].Count && __obf_59eea0bd834c59bb[__obf_c03cf6504cc3cf88].Key > __obf_59eea0bd834c59bb[__obf_c9b41611d2cb9abb].Key)
}

func (__obf_59eea0bd834c59bb Nodes) Swap(__obf_c03cf6504cc3cf88, __obf_c9b41611d2cb9abb int) {
	__obf_59eea0bd834c59bb[__obf_c03cf6504cc3cf88], __obf_59eea0bd834c59bb[__obf_c9b41611d2cb9abb] = __obf_59eea0bd834c59bb[__obf_c9b41611d2cb9abb], __obf_59eea0bd834c59bb[__obf_c03cf6504cc3cf88]
}

func (__obf_59eea0bd834c59bb *Nodes) Push(__obf_e209a94366a528e6 any) {
	*__obf_59eea0bd834c59bb = append(*__obf_59eea0bd834c59bb, __obf_e209a94366a528e6.(*Node))
}

func (__obf_59eea0bd834c59bb *Nodes) Pop() any {
	var __obf_e209a94366a528e6 *Node
	__obf_e209a94366a528e6, *__obf_59eea0bd834c59bb = (*__obf_59eea0bd834c59bb)[len((*__obf_59eea0bd834c59bb))-1], (*__obf_59eea0bd834c59bb)[:len((*__obf_59eea0bd834c59bb))-1]
	return __obf_e209a94366a528e6
}
