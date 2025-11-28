package __obf_62eba4024f8fa381

import (
	"container/heap"
	"sort"
)

type Heap struct {
	Nodes Nodes
	K     uint32
}

func NewHeap(__obf_f067a403d3b1b52a uint32) *Heap {
	__obf_1271301575c5c2f5 := Nodes{}
	heap.Init(&__obf_1271301575c5c2f5)
	return &Heap{Nodes: __obf_1271301575c5c2f5, K: __obf_f067a403d3b1b52a}
}

func (__obf_1271301575c5c2f5 *Heap) Add(__obf_ce934fb34aed9edb *Node) *Node {
	if __obf_1271301575c5c2f5.K > uint32(len(__obf_1271301575c5c2f5.Nodes)) {
		heap.Push(&__obf_1271301575c5c2f5.Nodes, __obf_ce934fb34aed9edb)
	} else if __obf_ce934fb34aed9edb.Count > __obf_1271301575c5c2f5.Nodes[0].Count {
		__obf_f4a99d02336ede45 := heap.Pop(&__obf_1271301575c5c2f5.Nodes)
		heap.Push(&__obf_1271301575c5c2f5.Nodes, __obf_ce934fb34aed9edb)
		__obf_fed93f71ba4da61d := __obf_f4a99d02336ede45.(*Node)
		return __obf_fed93f71ba4da61d
	}
	return nil
}

func (__obf_1271301575c5c2f5 *Heap) Pop() *Node {
	__obf_f4a99d02336ede45 := heap.Pop(&__obf_1271301575c5c2f5.Nodes)
	return __obf_f4a99d02336ede45.(*Node)
}

func (__obf_1271301575c5c2f5 *Heap) Fix(__obf_82e59fdf75a6183f int, __obf_6c61f1671393b933 uint32) {
	__obf_1271301575c5c2f5.Nodes[__obf_82e59fdf75a6183f].Count = __obf_6c61f1671393b933
	heap.Fix(&__obf_1271301575c5c2f5.Nodes, __obf_82e59fdf75a6183f)
}

func (__obf_1271301575c5c2f5 *Heap) Min() uint32 {
	if len(__obf_1271301575c5c2f5.Nodes) == 0 {
		return 0
	}
	return __obf_1271301575c5c2f5.Nodes[0].Count
}

func (__obf_1271301575c5c2f5 *Heap) Find(__obf_df070ab4c712506c string) (int, bool) {
	for __obf_1376e2b9d2469c18 := range __obf_1271301575c5c2f5.Nodes {
		if __obf_1271301575c5c2f5.Nodes[__obf_1376e2b9d2469c18].Key == __obf_df070ab4c712506c {
			return __obf_1376e2b9d2469c18, true
		}
	}
	return 0, false
}

func (__obf_1271301575c5c2f5 *Heap) Sorted() Nodes {
	__obf_92b53f6ff9abf638 := append([]*Node(nil), __obf_1271301575c5c2f5.Nodes...)
	sort.Sort(sort.Reverse(Nodes(__obf_92b53f6ff9abf638)))
	return __obf_92b53f6ff9abf638
}

type Nodes []*Node

type Node struct {
	Key   string
	Count uint32
}

func (__obf_409580b446b55be1 Nodes) Len() int {
	return len(__obf_409580b446b55be1)
}

func (__obf_409580b446b55be1 Nodes) Less(__obf_1376e2b9d2469c18, __obf_d605aaaaca8d9c7c int) bool {
	return (__obf_409580b446b55be1[__obf_1376e2b9d2469c18].Count < __obf_409580b446b55be1[__obf_d605aaaaca8d9c7c].Count) || (__obf_409580b446b55be1[__obf_1376e2b9d2469c18].Count == __obf_409580b446b55be1[__obf_d605aaaaca8d9c7c].Count && __obf_409580b446b55be1[__obf_1376e2b9d2469c18].Key > __obf_409580b446b55be1[__obf_d605aaaaca8d9c7c].Key)
}

func (__obf_409580b446b55be1 Nodes) Swap(__obf_1376e2b9d2469c18, __obf_d605aaaaca8d9c7c int) {
	__obf_409580b446b55be1[__obf_1376e2b9d2469c18], __obf_409580b446b55be1[__obf_d605aaaaca8d9c7c] = __obf_409580b446b55be1[__obf_d605aaaaca8d9c7c], __obf_409580b446b55be1[__obf_1376e2b9d2469c18]
}

func (__obf_409580b446b55be1 *Nodes) Push(__obf_ce934fb34aed9edb any) {
	*__obf_409580b446b55be1 = append(*__obf_409580b446b55be1, __obf_ce934fb34aed9edb.(*Node))
}

func (__obf_409580b446b55be1 *Nodes) Pop() any {
	var __obf_ce934fb34aed9edb *Node
	__obf_ce934fb34aed9edb, *__obf_409580b446b55be1 = (*__obf_409580b446b55be1)[len((*__obf_409580b446b55be1))-1], (*__obf_409580b446b55be1)[:len((*__obf_409580b446b55be1))-1]
	return __obf_ce934fb34aed9edb
}
