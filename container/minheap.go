package __obf_038560a94647875f

import (
	"container/heap"
	"sort"
)

type Heap struct {
	Nodes Nodes
	K     uint32
}

func NewHeap(__obf_ed002d86aefb8e41 uint32) *Heap {
	__obf_85e761a4baea4bf9 := Nodes{}
	heap.Init(&__obf_85e761a4baea4bf9)
	return &Heap{Nodes: __obf_85e761a4baea4bf9, K: __obf_ed002d86aefb8e41}
}

func (__obf_85e761a4baea4bf9 *Heap) Add(__obf_80765176f0eb947b *Node) *Node {
	if __obf_85e761a4baea4bf9.K > uint32(len(__obf_85e761a4baea4bf9.Nodes)) {
		heap.Push(&__obf_85e761a4baea4bf9.Nodes, __obf_80765176f0eb947b)
	} else if __obf_80765176f0eb947b.Count > __obf_85e761a4baea4bf9.Nodes[0].Count {
		__obf_54333d77e0021155 := heap.Pop(&__obf_85e761a4baea4bf9.Nodes)
		heap.Push(&__obf_85e761a4baea4bf9.Nodes, __obf_80765176f0eb947b)
		__obf_67e2dc1170e23d58 := __obf_54333d77e0021155.(*Node)
		return __obf_67e2dc1170e23d58
	}
	return nil
}

func (__obf_85e761a4baea4bf9 *Heap) Pop() *Node {
	__obf_54333d77e0021155 := heap.Pop(&__obf_85e761a4baea4bf9.Nodes)
	return __obf_54333d77e0021155.(*Node)
}

func (__obf_85e761a4baea4bf9 *Heap) Fix(__obf_610b93b6ca9b4f6c int, __obf_8b99337c97e3c5e6 uint32) {
	__obf_85e761a4baea4bf9.
		Nodes[__obf_610b93b6ca9b4f6c].Count = __obf_8b99337c97e3c5e6
	heap.Fix(&__obf_85e761a4baea4bf9.Nodes, __obf_610b93b6ca9b4f6c)
}

func (__obf_85e761a4baea4bf9 *Heap) Min() uint32 {
	if len(__obf_85e761a4baea4bf9.Nodes) == 0 {
		return 0
	}
	return __obf_85e761a4baea4bf9.Nodes[0].Count
}

func (__obf_85e761a4baea4bf9 *Heap) Find(__obf_52e73bb48e810dd2 string) (int, bool) {
	for __obf_d2ea251c77215677 := range __obf_85e761a4baea4bf9.Nodes {
		if __obf_85e761a4baea4bf9.Nodes[__obf_d2ea251c77215677].Key == __obf_52e73bb48e810dd2 {
			return __obf_d2ea251c77215677, true
		}
	}
	return 0, false
}

func (__obf_85e761a4baea4bf9 *Heap) Sorted() Nodes {
	__obf_884f77096b1442d3 := append([]*Node(nil), __obf_85e761a4baea4bf9.Nodes...)
	sort.Sort(sort.Reverse(Nodes(__obf_884f77096b1442d3)))
	return __obf_884f77096b1442d3
}

type Nodes []*Node

type Node struct {
	Key   string
	Count uint32
}

func (__obf_406632683c089fe0 Nodes) Len() int {
	return len(__obf_406632683c089fe0)
}

func (__obf_406632683c089fe0 Nodes) Less(__obf_d2ea251c77215677, __obf_10f61d83892390d8 int) bool {
	return (__obf_406632683c089fe0[__obf_d2ea251c77215677].Count < __obf_406632683c089fe0[__obf_10f61d83892390d8].Count) || (__obf_406632683c089fe0[__obf_d2ea251c77215677].Count == __obf_406632683c089fe0[__obf_10f61d83892390d8].Count && __obf_406632683c089fe0[__obf_d2ea251c77215677].Key > __obf_406632683c089fe0[__obf_10f61d83892390d8].Key)
}

func (__obf_406632683c089fe0 Nodes) Swap(__obf_d2ea251c77215677, __obf_10f61d83892390d8 int) {
	__obf_406632683c089fe0[__obf_d2ea251c77215677], __obf_406632683c089fe0[__obf_10f61d83892390d8] = __obf_406632683c089fe0[__obf_10f61d83892390d8], __obf_406632683c089fe0[__obf_d2ea251c77215677]
}

func (__obf_406632683c089fe0 *Nodes) Push(__obf_80765176f0eb947b any) {
	*__obf_406632683c089fe0 = append(*__obf_406632683c089fe0, __obf_80765176f0eb947b.(*Node))
}

func (__obf_406632683c089fe0 *Nodes) Pop() any {
	var __obf_80765176f0eb947b *Node
	__obf_80765176f0eb947b, *__obf_406632683c089fe0 = (*__obf_406632683c089fe0)[len((*__obf_406632683c089fe0))-1], (*__obf_406632683c089fe0)[:len((*__obf_406632683c089fe0))-1]
	return __obf_80765176f0eb947b
}
