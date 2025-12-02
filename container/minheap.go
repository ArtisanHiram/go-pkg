package __obf_4a16ef421fb74992

import (
	"container/heap"
	"sort"
)

type Heap struct {
	Nodes Nodes
	K     uint32
}

func NewHeap(__obf_91fb3c93595a08ac uint32) *Heap {
	__obf_55fe93b5d3e19f79 := Nodes{}
	heap.Init(&__obf_55fe93b5d3e19f79)
	return &Heap{Nodes: __obf_55fe93b5d3e19f79, K: __obf_91fb3c93595a08ac}
}

func (__obf_55fe93b5d3e19f79 *Heap) Add(__obf_db7248ecee87d860 *Node) *Node {
	if __obf_55fe93b5d3e19f79.K > uint32(len(__obf_55fe93b5d3e19f79.Nodes)) {
		heap.Push(&__obf_55fe93b5d3e19f79.Nodes, __obf_db7248ecee87d860)
	} else if __obf_db7248ecee87d860.Count > __obf_55fe93b5d3e19f79.Nodes[0].Count {
		__obf_db9b39bdb3b44923 := heap.Pop(&__obf_55fe93b5d3e19f79.Nodes)
		heap.Push(&__obf_55fe93b5d3e19f79.Nodes, __obf_db7248ecee87d860)
		__obf_cb94efc111ad4220 := __obf_db9b39bdb3b44923.(*Node)
		return __obf_cb94efc111ad4220
	}
	return nil
}

func (__obf_55fe93b5d3e19f79 *Heap) Pop() *Node {
	__obf_db9b39bdb3b44923 := heap.Pop(&__obf_55fe93b5d3e19f79.Nodes)
	return __obf_db9b39bdb3b44923.(*Node)
}

func (__obf_55fe93b5d3e19f79 *Heap) Fix(__obf_a0811e0731ad212a int, __obf_1d7629104d39aeed uint32) {
	__obf_55fe93b5d3e19f79.
		Nodes[__obf_a0811e0731ad212a].Count = __obf_1d7629104d39aeed
	heap.Fix(&__obf_55fe93b5d3e19f79.Nodes, __obf_a0811e0731ad212a)
}

func (__obf_55fe93b5d3e19f79 *Heap) Min() uint32 {
	if len(__obf_55fe93b5d3e19f79.Nodes) == 0 {
		return 0
	}
	return __obf_55fe93b5d3e19f79.Nodes[0].Count
}

func (__obf_55fe93b5d3e19f79 *Heap) Find(__obf_d1a1f98ae0fb119e string) (int, bool) {
	for __obf_79cf1c44489eaf01 := range __obf_55fe93b5d3e19f79.Nodes {
		if __obf_55fe93b5d3e19f79.Nodes[__obf_79cf1c44489eaf01].Key == __obf_d1a1f98ae0fb119e {
			return __obf_79cf1c44489eaf01, true
		}
	}
	return 0, false
}

func (__obf_55fe93b5d3e19f79 *Heap) Sorted() Nodes {
	__obf_aad976bcca520eb3 := append([]*Node(nil), __obf_55fe93b5d3e19f79.Nodes...)
	sort.Sort(sort.Reverse(Nodes(__obf_aad976bcca520eb3)))
	return __obf_aad976bcca520eb3
}

type Nodes []*Node

type Node struct {
	Key   string
	Count uint32
}

func (__obf_41509bed52f90987 Nodes) Len() int {
	return len(__obf_41509bed52f90987)
}

func (__obf_41509bed52f90987 Nodes) Less(__obf_79cf1c44489eaf01, __obf_aed15258f3705fd9 int) bool {
	return (__obf_41509bed52f90987[__obf_79cf1c44489eaf01].Count < __obf_41509bed52f90987[__obf_aed15258f3705fd9].Count) || (__obf_41509bed52f90987[__obf_79cf1c44489eaf01].Count == __obf_41509bed52f90987[__obf_aed15258f3705fd9].Count && __obf_41509bed52f90987[__obf_79cf1c44489eaf01].Key > __obf_41509bed52f90987[__obf_aed15258f3705fd9].Key)
}

func (__obf_41509bed52f90987 Nodes) Swap(__obf_79cf1c44489eaf01, __obf_aed15258f3705fd9 int) {
	__obf_41509bed52f90987[__obf_79cf1c44489eaf01], __obf_41509bed52f90987[__obf_aed15258f3705fd9] = __obf_41509bed52f90987[__obf_aed15258f3705fd9], __obf_41509bed52f90987[__obf_79cf1c44489eaf01]
}

func (__obf_41509bed52f90987 *Nodes) Push(__obf_db7248ecee87d860 any) {
	*__obf_41509bed52f90987 = append(*__obf_41509bed52f90987, __obf_db7248ecee87d860.(*Node))
}

func (__obf_41509bed52f90987 *Nodes) Pop() any {
	var __obf_db7248ecee87d860 *Node
	__obf_db7248ecee87d860, *__obf_41509bed52f90987 = (*__obf_41509bed52f90987)[len((*__obf_41509bed52f90987))-1], (*__obf_41509bed52f90987)[:len((*__obf_41509bed52f90987))-1]
	return __obf_db7248ecee87d860
}
