package __obf_9004b47202f9204b

import (
	"container/heap"
	"sort"
)

type Heap struct {
	Nodes Nodes
	K     uint32
}

func NewHeap(__obf_cc18ca82ac13013e uint32) *Heap {
	__obf_e64b3bb2d46916d3 := Nodes{}
	heap.Init(&__obf_e64b3bb2d46916d3)
	return &Heap{Nodes: __obf_e64b3bb2d46916d3, K: __obf_cc18ca82ac13013e}
}

func (__obf_e64b3bb2d46916d3 *Heap) Add(__obf_903725be6b5384b7 *Node) *Node {
	if __obf_e64b3bb2d46916d3.K > uint32(len(__obf_e64b3bb2d46916d3.Nodes)) {
		heap.Push(&__obf_e64b3bb2d46916d3.Nodes, __obf_903725be6b5384b7)
	} else if __obf_903725be6b5384b7.Count > __obf_e64b3bb2d46916d3.Nodes[0].Count {
		__obf_f64984ee664470d5 := heap.Pop(&__obf_e64b3bb2d46916d3.Nodes)
		heap.Push(&__obf_e64b3bb2d46916d3.Nodes, __obf_903725be6b5384b7)
		__obf_c0dad0405e4a8dfe := __obf_f64984ee664470d5.(*Node)
		return __obf_c0dad0405e4a8dfe
	}
	return nil
}

func (__obf_e64b3bb2d46916d3 *Heap) Pop() *Node {
	__obf_f64984ee664470d5 := heap.Pop(&__obf_e64b3bb2d46916d3.Nodes)
	return __obf_f64984ee664470d5.(*Node)
}

func (__obf_e64b3bb2d46916d3 *Heap) Fix(__obf_095e96d6bfc7baf7 int, __obf_4dd2f2e7ef923bcc uint32) {
	__obf_e64b3bb2d46916d3.
		Nodes[__obf_095e96d6bfc7baf7].Count = __obf_4dd2f2e7ef923bcc
	heap.Fix(&__obf_e64b3bb2d46916d3.Nodes, __obf_095e96d6bfc7baf7)
}

func (__obf_e64b3bb2d46916d3 *Heap) Min() uint32 {
	if len(__obf_e64b3bb2d46916d3.Nodes) == 0 {
		return 0
	}
	return __obf_e64b3bb2d46916d3.Nodes[0].Count
}

func (__obf_e64b3bb2d46916d3 *Heap) Find(__obf_355e7c922e433678 string) (int, bool) {
	for __obf_c97cf0c2d79c5ab6 := range __obf_e64b3bb2d46916d3.Nodes {
		if __obf_e64b3bb2d46916d3.Nodes[__obf_c97cf0c2d79c5ab6].Key == __obf_355e7c922e433678 {
			return __obf_c97cf0c2d79c5ab6, true
		}
	}
	return 0, false
}

func (__obf_e64b3bb2d46916d3 *Heap) Sorted() Nodes {
	__obf_2395646c7467f8c0 := append([]*Node(nil), __obf_e64b3bb2d46916d3.Nodes...)
	sort.Sort(sort.Reverse(Nodes(__obf_2395646c7467f8c0)))
	return __obf_2395646c7467f8c0
}

type Nodes []*Node

type Node struct {
	Key   string
	Count uint32
}

func (__obf_d2d1bdfa577ba7e8 Nodes) Len() int {
	return len(__obf_d2d1bdfa577ba7e8)
}

func (__obf_d2d1bdfa577ba7e8 Nodes) Less(__obf_c97cf0c2d79c5ab6, __obf_e0db4519fdc27298 int) bool {
	return (__obf_d2d1bdfa577ba7e8[__obf_c97cf0c2d79c5ab6].Count < __obf_d2d1bdfa577ba7e8[__obf_e0db4519fdc27298].Count) || (__obf_d2d1bdfa577ba7e8[__obf_c97cf0c2d79c5ab6].Count == __obf_d2d1bdfa577ba7e8[__obf_e0db4519fdc27298].Count && __obf_d2d1bdfa577ba7e8[__obf_c97cf0c2d79c5ab6].Key > __obf_d2d1bdfa577ba7e8[__obf_e0db4519fdc27298].Key)
}

func (__obf_d2d1bdfa577ba7e8 Nodes) Swap(__obf_c97cf0c2d79c5ab6, __obf_e0db4519fdc27298 int) {
	__obf_d2d1bdfa577ba7e8[__obf_c97cf0c2d79c5ab6], __obf_d2d1bdfa577ba7e8[__obf_e0db4519fdc27298] = __obf_d2d1bdfa577ba7e8[__obf_e0db4519fdc27298], __obf_d2d1bdfa577ba7e8[__obf_c97cf0c2d79c5ab6]
}

func (__obf_d2d1bdfa577ba7e8 *Nodes) Push(__obf_903725be6b5384b7 any) {
	*__obf_d2d1bdfa577ba7e8 = append(*__obf_d2d1bdfa577ba7e8, __obf_903725be6b5384b7.(*Node))
}

func (__obf_d2d1bdfa577ba7e8 *Nodes) Pop() any {
	var __obf_903725be6b5384b7 *Node
	__obf_903725be6b5384b7, *__obf_d2d1bdfa577ba7e8 = (*__obf_d2d1bdfa577ba7e8)[len((*__obf_d2d1bdfa577ba7e8))-1], (*__obf_d2d1bdfa577ba7e8)[:len((*__obf_d2d1bdfa577ba7e8))-1]
	return __obf_903725be6b5384b7
}
