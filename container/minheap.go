package __obf_76599ab96a208083

import (
	"container/heap"
	"sort"
)

type Heap struct {
	Nodes Nodes
	K     uint32
}

func NewHeap(__obf_2971938b3f770fd1 uint32) *Heap {
	__obf_703df7aac448a8d0 := Nodes{}
	heap.Init(&__obf_703df7aac448a8d0)
	return &Heap{Nodes: __obf_703df7aac448a8d0, K: __obf_2971938b3f770fd1}
}

func (__obf_703df7aac448a8d0 *Heap) Add(__obf_abdb95bb10f251fe *Node) *Node {
	if __obf_703df7aac448a8d0.K > uint32(len(__obf_703df7aac448a8d0.Nodes)) {
		heap.Push(&__obf_703df7aac448a8d0.Nodes, __obf_abdb95bb10f251fe)
	} else if __obf_abdb95bb10f251fe.Count > __obf_703df7aac448a8d0.Nodes[0].Count {
		__obf_6469c6f8bc8923a1 := heap.Pop(&__obf_703df7aac448a8d0.Nodes)
		heap.Push(&__obf_703df7aac448a8d0.Nodes, __obf_abdb95bb10f251fe)
		__obf_ee28ba4e9ac049f8 := __obf_6469c6f8bc8923a1.(*Node)
		return __obf_ee28ba4e9ac049f8
	}
	return nil
}

func (__obf_703df7aac448a8d0 *Heap) Pop() *Node {
	__obf_6469c6f8bc8923a1 := heap.Pop(&__obf_703df7aac448a8d0.Nodes)
	return __obf_6469c6f8bc8923a1.(*Node)
}

func (__obf_703df7aac448a8d0 *Heap) Fix(__obf_635d4198a10d544b int, __obf_c5d600e67da7e0a9 uint32) {
	__obf_703df7aac448a8d0.Nodes[__obf_635d4198a10d544b].Count = __obf_c5d600e67da7e0a9
	heap.Fix(&__obf_703df7aac448a8d0.Nodes, __obf_635d4198a10d544b)
}

func (__obf_703df7aac448a8d0 *Heap) Min() uint32 {
	if len(__obf_703df7aac448a8d0.Nodes) == 0 {
		return 0
	}
	return __obf_703df7aac448a8d0.Nodes[0].Count
}

func (__obf_703df7aac448a8d0 *Heap) Find(__obf_4368715d0a6d4f0b string) (int, bool) {
	for __obf_8d4edb68382a63c4 := range __obf_703df7aac448a8d0.Nodes {
		if __obf_703df7aac448a8d0.Nodes[__obf_8d4edb68382a63c4].Key == __obf_4368715d0a6d4f0b {
			return __obf_8d4edb68382a63c4, true
		}
	}
	return 0, false
}

func (__obf_703df7aac448a8d0 *Heap) Sorted() Nodes {
	__obf_22d5f491c1dcf38b := append([]*Node(nil), __obf_703df7aac448a8d0.Nodes...)
	sort.Sort(sort.Reverse(Nodes(__obf_22d5f491c1dcf38b)))
	return __obf_22d5f491c1dcf38b
}

type Nodes []*Node

type Node struct {
	Key   string
	Count uint32
}

func (__obf_18a54e2632d73807 Nodes) Len() int {
	return len(__obf_18a54e2632d73807)
}

func (__obf_18a54e2632d73807 Nodes) Less(__obf_8d4edb68382a63c4, __obf_5c6ea4abd253b60e int) bool {
	return (__obf_18a54e2632d73807[__obf_8d4edb68382a63c4].Count < __obf_18a54e2632d73807[__obf_5c6ea4abd253b60e].Count) || (__obf_18a54e2632d73807[__obf_8d4edb68382a63c4].Count == __obf_18a54e2632d73807[__obf_5c6ea4abd253b60e].Count && __obf_18a54e2632d73807[__obf_8d4edb68382a63c4].Key > __obf_18a54e2632d73807[__obf_5c6ea4abd253b60e].Key)
}

func (__obf_18a54e2632d73807 Nodes) Swap(__obf_8d4edb68382a63c4, __obf_5c6ea4abd253b60e int) {
	__obf_18a54e2632d73807[__obf_8d4edb68382a63c4], __obf_18a54e2632d73807[__obf_5c6ea4abd253b60e] = __obf_18a54e2632d73807[__obf_5c6ea4abd253b60e], __obf_18a54e2632d73807[__obf_8d4edb68382a63c4]
}

func (__obf_18a54e2632d73807 *Nodes) Push(__obf_abdb95bb10f251fe any) {
	*__obf_18a54e2632d73807 = append(*__obf_18a54e2632d73807, __obf_abdb95bb10f251fe.(*Node))
}

func (__obf_18a54e2632d73807 *Nodes) Pop() any {
	var __obf_abdb95bb10f251fe *Node
	__obf_abdb95bb10f251fe, *__obf_18a54e2632d73807 = (*__obf_18a54e2632d73807)[len((*__obf_18a54e2632d73807))-1], (*__obf_18a54e2632d73807)[:len((*__obf_18a54e2632d73807))-1]
	return __obf_abdb95bb10f251fe
}
