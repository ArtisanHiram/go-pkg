package __obf_9861fa13140c30a3

import (
	"container/heap"
	"sort"
)

type Heap struct {
	Nodes Nodes
	K     uint32
}

func NewHeap(__obf_299b9ac8f6041f1b uint32) *Heap {
	__obf_aceb1679325404db := Nodes{}
	heap.Init(&__obf_aceb1679325404db)
	return &Heap{Nodes: __obf_aceb1679325404db, K: __obf_299b9ac8f6041f1b}
}

func (__obf_aceb1679325404db *Heap) Add(__obf_bbf5522c3a9df879 *Node) *Node {
	if __obf_aceb1679325404db.K > uint32(len(__obf_aceb1679325404db.Nodes)) {
		heap.Push(&__obf_aceb1679325404db.Nodes, __obf_bbf5522c3a9df879)
	} else if __obf_bbf5522c3a9df879.Count > __obf_aceb1679325404db.Nodes[0].Count {
		__obf_957ce391d77b695f := heap.Pop(&__obf_aceb1679325404db.Nodes)
		heap.Push(&__obf_aceb1679325404db.Nodes, __obf_bbf5522c3a9df879)
		__obf_2b073328ceed4884 := __obf_957ce391d77b695f.(*Node)
		return __obf_2b073328ceed4884
	}
	return nil
}

func (__obf_aceb1679325404db *Heap) Pop() *Node {
	__obf_957ce391d77b695f := heap.Pop(&__obf_aceb1679325404db.Nodes)
	return __obf_957ce391d77b695f.(*Node)
}

func (__obf_aceb1679325404db *Heap) Fix(__obf_55d78e2ec3bee11d int, __obf_7e6767c22d23f06d uint32) {
	__obf_aceb1679325404db.Nodes[__obf_55d78e2ec3bee11d].Count = __obf_7e6767c22d23f06d
	heap.Fix(&__obf_aceb1679325404db.Nodes, __obf_55d78e2ec3bee11d)
}

func (__obf_aceb1679325404db *Heap) Min() uint32 {
	if len(__obf_aceb1679325404db.Nodes) == 0 {
		return 0
	}
	return __obf_aceb1679325404db.Nodes[0].Count
}

func (__obf_aceb1679325404db *Heap) Find(__obf_43194ec765d86867 string) (int, bool) {
	for __obf_35fc60e92e4b05ba := range __obf_aceb1679325404db.Nodes {
		if __obf_aceb1679325404db.Nodes[__obf_35fc60e92e4b05ba].Key == __obf_43194ec765d86867 {
			return __obf_35fc60e92e4b05ba, true
		}
	}
	return 0, false
}

func (__obf_aceb1679325404db *Heap) Sorted() Nodes {
	__obf_5eab0c876476504e := append([]*Node(nil), __obf_aceb1679325404db.Nodes...)
	sort.Sort(sort.Reverse(Nodes(__obf_5eab0c876476504e)))
	return __obf_5eab0c876476504e
}

type Nodes []*Node

type Node struct {
	Key   string
	Count uint32
}

func (__obf_bdf68a51ae002dc2 Nodes) Len() int {
	return len(__obf_bdf68a51ae002dc2)
}

func (__obf_bdf68a51ae002dc2 Nodes) Less(__obf_35fc60e92e4b05ba, __obf_ebe85d0a1d7e1248 int) bool {
	return (__obf_bdf68a51ae002dc2[__obf_35fc60e92e4b05ba].Count < __obf_bdf68a51ae002dc2[__obf_ebe85d0a1d7e1248].Count) || (__obf_bdf68a51ae002dc2[__obf_35fc60e92e4b05ba].Count == __obf_bdf68a51ae002dc2[__obf_ebe85d0a1d7e1248].Count && __obf_bdf68a51ae002dc2[__obf_35fc60e92e4b05ba].Key > __obf_bdf68a51ae002dc2[__obf_ebe85d0a1d7e1248].Key)
}

func (__obf_bdf68a51ae002dc2 Nodes) Swap(__obf_35fc60e92e4b05ba, __obf_ebe85d0a1d7e1248 int) {
	__obf_bdf68a51ae002dc2[__obf_35fc60e92e4b05ba], __obf_bdf68a51ae002dc2[__obf_ebe85d0a1d7e1248] = __obf_bdf68a51ae002dc2[__obf_ebe85d0a1d7e1248], __obf_bdf68a51ae002dc2[__obf_35fc60e92e4b05ba]
}

func (__obf_bdf68a51ae002dc2 *Nodes) Push(__obf_bbf5522c3a9df879 any) {
	*__obf_bdf68a51ae002dc2 = append(*__obf_bdf68a51ae002dc2, __obf_bbf5522c3a9df879.(*Node))
}

func (__obf_bdf68a51ae002dc2 *Nodes) Pop() any {
	var __obf_bbf5522c3a9df879 *Node
	__obf_bbf5522c3a9df879, *__obf_bdf68a51ae002dc2 = (*__obf_bdf68a51ae002dc2)[len((*__obf_bdf68a51ae002dc2))-1], (*__obf_bdf68a51ae002dc2)[:len((*__obf_bdf68a51ae002dc2))-1]
	return __obf_bbf5522c3a9df879
}
