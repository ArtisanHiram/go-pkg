package __obf_b0bebe5eb45b8ad6

import (
	"container/heap"
	"sort"
)

type Heap struct {
	Nodes Nodes
	K     uint32
}

func NewHeap(__obf_dca1c5db9aad12e8 uint32) *Heap {
	__obf_98871087bafa5bd1 := Nodes{}
	heap.Init(&__obf_98871087bafa5bd1)
	return &Heap{Nodes: __obf_98871087bafa5bd1, K: __obf_dca1c5db9aad12e8}
}

func (__obf_98871087bafa5bd1 *Heap) Add(__obf_c16b641e2cd6ad69 *Node) *Node {
	if __obf_98871087bafa5bd1.K > uint32(len(__obf_98871087bafa5bd1.Nodes)) {
		heap.Push(&__obf_98871087bafa5bd1.Nodes, __obf_c16b641e2cd6ad69)
	} else if __obf_c16b641e2cd6ad69.Count > __obf_98871087bafa5bd1.Nodes[0].Count {
		__obf_891ef6f642020acf := heap.Pop(&__obf_98871087bafa5bd1.Nodes)
		heap.Push(&__obf_98871087bafa5bd1.Nodes, __obf_c16b641e2cd6ad69)
		__obf_612cd42ce34847be := __obf_891ef6f642020acf.(*Node)
		return __obf_612cd42ce34847be
	}
	return nil
}

func (__obf_98871087bafa5bd1 *Heap) Pop() *Node {
	__obf_891ef6f642020acf := heap.Pop(&__obf_98871087bafa5bd1.Nodes)
	return __obf_891ef6f642020acf.(*Node)
}

func (__obf_98871087bafa5bd1 *Heap) Fix(__obf_12d332fd9cbd0f8d int, __obf_39f45b145b5eac3d uint32) {
	__obf_98871087bafa5bd1.
		Nodes[__obf_12d332fd9cbd0f8d].Count = __obf_39f45b145b5eac3d
	heap.Fix(&__obf_98871087bafa5bd1.Nodes, __obf_12d332fd9cbd0f8d)
}

func (__obf_98871087bafa5bd1 *Heap) Min() uint32 {
	if len(__obf_98871087bafa5bd1.Nodes) == 0 {
		return 0
	}
	return __obf_98871087bafa5bd1.Nodes[0].Count
}

func (__obf_98871087bafa5bd1 *Heap) Find(__obf_9c15798bcb95be3e string) (int, bool) {
	for __obf_095a614fd95fc26e := range __obf_98871087bafa5bd1.Nodes {
		if __obf_98871087bafa5bd1.Nodes[__obf_095a614fd95fc26e].Key == __obf_9c15798bcb95be3e {
			return __obf_095a614fd95fc26e, true
		}
	}
	return 0, false
}

func (__obf_98871087bafa5bd1 *Heap) Sorted() Nodes {
	__obf_f171fa45468e7200 := append([]*Node(nil), __obf_98871087bafa5bd1.Nodes...)
	sort.Sort(sort.Reverse(Nodes(__obf_f171fa45468e7200)))
	return __obf_f171fa45468e7200
}

type Nodes []*Node

type Node struct {
	Key   string
	Count uint32
}

func (__obf_fc5555a986252f47 Nodes) Len() int {
	return len(__obf_fc5555a986252f47)
}

func (__obf_fc5555a986252f47 Nodes) Less(__obf_095a614fd95fc26e, __obf_017d2353b8e02ea9 int) bool {
	return (__obf_fc5555a986252f47[__obf_095a614fd95fc26e].Count < __obf_fc5555a986252f47[__obf_017d2353b8e02ea9].Count) || (__obf_fc5555a986252f47[__obf_095a614fd95fc26e].Count == __obf_fc5555a986252f47[__obf_017d2353b8e02ea9].Count && __obf_fc5555a986252f47[__obf_095a614fd95fc26e].Key > __obf_fc5555a986252f47[__obf_017d2353b8e02ea9].Key)
}

func (__obf_fc5555a986252f47 Nodes) Swap(__obf_095a614fd95fc26e, __obf_017d2353b8e02ea9 int) {
	__obf_fc5555a986252f47[__obf_095a614fd95fc26e], __obf_fc5555a986252f47[__obf_017d2353b8e02ea9] = __obf_fc5555a986252f47[__obf_017d2353b8e02ea9], __obf_fc5555a986252f47[__obf_095a614fd95fc26e]
}

func (__obf_fc5555a986252f47 *Nodes) Push(__obf_c16b641e2cd6ad69 any) {
	*__obf_fc5555a986252f47 = append(*__obf_fc5555a986252f47, __obf_c16b641e2cd6ad69.(*Node))
}

func (__obf_fc5555a986252f47 *Nodes) Pop() any {
	var __obf_c16b641e2cd6ad69 *Node
	__obf_c16b641e2cd6ad69, *__obf_fc5555a986252f47 = (*__obf_fc5555a986252f47)[len((*__obf_fc5555a986252f47))-1], (*__obf_fc5555a986252f47)[:len((*__obf_fc5555a986252f47))-1]
	return __obf_c16b641e2cd6ad69
}
